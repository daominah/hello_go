package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime)

	// the following values are configured in Keycloak:

	// Keycloak:22.0.5 run in Docker
	kcAddress := "http://localhost:18080"
	kcRealm := "myrealm"
	kcClientID := "myclient"
	kcClientSecret := "BE2E37zif7WTPq1eiznToeCuVHHmAW1L"
	kcRedirect := "http://localhost:8181/valid-redirect-uri-after-login-keycloak" // redirect browser here after a successful login

	// https://stackoverflow.com/questions/48855122/keycloak-adaptor-for-golang-application:

	issuer := fmt.Sprintf("%v/realms/%v", kcAddress, kcRealm)
	provider, err := oidc.NewProvider(context.TODO(), issuer)
	if err != nil {
		panic(err)
	}

	oauth2Config := oauth2.Config{
		ClientID:     kcClientID,
		ClientSecret: kcClientSecret,
		RedirectURL:  kcRedirect,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
	// Endpoint.AuthURL = realms/myrealm/protocol/openid-connect/auth

	// SkipClientIDCheck bacause Keycloak token has "azp" field,
	// OIDC specs about "azp", "aud" are confusing so just skip
	verifier := provider.Verifier(&oidc.Config{SkipClientIDCheck: true})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rawAccessToken := r.Header.Get("Authorization")
		parts := strings.Split(rawAccessToken, " ")
		if len(parts) != 2 {
			log.Printf("Authorization does not follow format `{token_type} {access_token}`, redirect to Keycloak login page")
			state := randomString()
			http.SetCookie(w, &http.Cookie{
				Name:     StateCookieName,
				Value:    state,
				MaxAge:   int(time.Hour.Seconds()),
				HttpOnly: true, // HttpOnly cookie cannot be read by JS
				Secure:   r.TLS != nil,
			})
			http.Redirect(w, r, oauth2Config.AuthCodeURL(state), http.StatusFound)
			return
		}
		accessToken := parts[1]
		token, err := verifier.Verify(context.TODO(), accessToken)
		if err != nil {
			errMsg := fmt.Sprintf("error verifier.Verify: %v, %+v", err, token)
			log.Printf(errMsg)
			http.Error(w, errMsg, http.StatusUnauthorized)
		}
		_ = token
		w.Write([]byte("hello, you are logged in successfully"))
	})

	http.HandleFunc("/valid-redirect-uri-after-login-keycloak", func(w http.ResponseWriter, r *http.Request) {
		state, err := r.Cookie(StateCookieName)
		if err != nil {
			http.Error(w, "state not found", http.StatusBadRequest)
			return
		}
		if r.FormValue("state") != state.Value {
			errMsg := fmt.Sprintf("state not matched")
			//log.Printf("state not matched: url.state: %v, cookie.state: %v", r.FormValue("state"), state.Value)
			http.Error(w, errMsg, http.StatusBadRequest)
			return
		}

		oauth2Token, err := oauth2Config.Exchange(context.TODO(), r.FormValue("code"))
		if err != nil {
			errMsg := fmt.Sprintf("error oauth2Config.Exchange: %v", err)
			log.Printf(errMsg)
			http.Error(w, errMsg, http.StatusInternalServerError)
			return
		}

		idTokenStr, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			errMsg := fmt.Sprintf("error oauth2Token.Extra: %v", err)
			log.Printf(errMsg)
			http.Error(w, errMsg, http.StatusInternalServerError)
			return
		}
		//log.Printf("idTokenStr: %v", idTokenStr)
		idToken, err := verifier.Verify(context.TODO(), idTokenStr)
		if err != nil {
			errMsg := fmt.Sprintf("error verifier.Verify: %v", err)
			log.Printf(errMsg)
			http.Error(w, errMsg, http.StatusInternalServerError)
			return
		}
		var idTokenDetail map[string]any
		err = idToken.Claims(&idTokenDetail)
		if err != nil {
			errMsg := fmt.Sprintf("error idToken.Claims: %v", err)
			log.Printf(errMsg)
			http.Error(w, errMsg, http.StatusInternalServerError)
			return
		}
		userEmail, ok := idTokenDetail["email"].(string)
		if !ok {
			errMsg := fmt.Sprintf("error missing 'email' in idToken.Claims: %#v", idTokenDetail["email"])
			log.Printf(errMsg)
			http.Error(w, errMsg, http.StatusInternalServerError)
		}
		log.Printf("userEmail: %v", userEmail)
		resp := struct {
			OAuth2Token   *oauth2.Token
			IDTokenClaims map[string]any
		}{
			OAuth2Token:   oauth2Token,
			IDTokenClaims: idTokenDetail,
		}
		beauty, err := json.MarshalIndent(resp, "", "\t")
		if err != nil {
			errMsg := fmt.Sprintf("error json.MarshalIndent: %v", err)
			log.Printf(errMsg)
			http.Error(w, errMsg, http.StatusInternalServerError)
			return
		}
		w.Write(beauty)
	})

	listenPort := ":8181"
	log.Printf("http://localhost%v", listenPort)
	err = http.ListenAndServe(listenPort, nil)
	log.Fatal("error ListenAndServe:", err)

	// export at=
	// curl -i -H "Authorization: Bearer ${at}" 'http://localhost:8181'
}

func randomString() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	r := hex.EncodeToString(b)
	return r
}

const StateCookieName = "state_daominah"
