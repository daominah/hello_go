package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
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

	issuer := fmt.Sprintf("%v/realms/%v", kcAddress, kcRealm)
	oidcProvider, err := oidc.NewProvider(context.TODO(), issuer)
	if err != nil {
		panic(err)
	}

	tseScopes := strings.Split("openid email profile get:my-org:all create:bases:all", " ")
	//tseScopes := strings.Split("openid email profile create:bases:all get:my-org-type:base put:my-org-type:base post:my-org-type-base:steps list:my-org-type-base:steps delete:my-org-type-base:steps list:my-org-type:base list:type-all:base post:my-org-type:base create:cases:all list:my-org-type-base:case get:my-org-type-base:case patch:my-org-type-base:case post:my-org-type-base:case get:my-org-type-base:analytic list:all:org list:my-org-type-base:notification create:users:all get:notification:all create:my-org:user list:my-org:type update:my-org:user get:types:all list:my-org:user list:all:type update:me:passwords create:types:all get:my-org-type-base:analytic-case list:my-org:tag post:my-org:tag get:my-org-type-base:analyticbyfield get:my-org-type-base:qualitative-feedback get:my-org:all put:my-org-type-base:case get:my-org-type-base:qualitative-feedback-all get:my-org:activity-report get:my-org:tag get:my-org:user post:case-delta:validation publish:my-org-type:base get:my-org-type-base:workflow post:my-org-type-base:workflow put:my-org-type-base:workflow delete:my-org-type-base:workflow post:my-org-type-base-workflow:action put:my-org-type-base-workflow:action list:my-org-type-base:workflow delete:my-org-type-base-workflow:action get:my-org-type-base-case:sync-status get:my-org-type-base:sync-status get:case:pass put:my-org-type:base-preferences", " ")
	_ = tseScopes
	oauth2Config := oauth2.Config{
		ClientID:     kcClientID,
		ClientSecret: kcClientSecret,
		RedirectURL:  kcRedirect,
		Endpoint:     oidcProvider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
		//Scopes:       tseScopes,
	}

	// SkipClientIDCheck bacause Keycloak token has "azp" field,
	// OIDC specs about "azp", "aud" are confusing so skip for now
	verifier := oidcProvider.Verifier(&oidc.Config{SkipClientIDCheck: true})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rawAccessToken := r.Header.Get("Authorization")
		parts := strings.Split(rawAccessToken, " ")
		if len(parts) != 2 {
			log.Printf("Authorization does not follow format `{token_type} {access_token}`, maybe not logged in, redirect to Keycloak login page")
			state := randomString(16)
			http.SetCookie(w, &http.Cookie{
				Name:     StateCookieName,
				Value:    state,
				MaxAge:   int(time.Hour.Seconds()),
				HttpOnly: true, // HttpOnly cookie cannot be read by JS
				Secure:   r.TLS != nil,
			})
			codeFlowURL := oauth2Config.AuthCodeURL(state)
			//log.Printf("authCodeURL: %v", codeFlowURL)
			tokenFlowURL := AuthImplicitFlowURL(oauth2Config, state)
			log.Printf("tokenFlowURL: %v", tokenFlowURL)
			_, _ = codeFlowURL, tokenFlowURL
			loginURL := codeFlowURL
			http.Redirect(w, r, loginURL, http.StatusFound)
			return
		}
		accessToken := parts[1]
		idToken, err := verifier.Verify(context.TODO(), accessToken)
		if err != nil {
			errMsg := fmt.Sprintf("error verifier.Verify: %v", err)
			log.Printf(errMsg)
			http.Error(w, errMsg, http.StatusUnauthorized)
			return
		}
		user, _ := ParseIDToken(idToken)
		log.Printf("user: %+v, idToken: %+v", user, ParseFullIDToken(idToken))
		w.Write([]byte(fmt.Sprintf("hello %v, you are logged in successfully", user.Email)))
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

		if r.FormValue("error") != "" {
			errMsg := fmt.Sprintf("error: %v\nerror_description:%v", r.FormValue("error"), r.FormValue("error_description"))
			http.Error(w, errMsg, http.StatusInternalServerError)
			return
		}

		oidcCode := r.FormValue("code")
		oauth2Token, err := oauth2Config.Exchange(context.TODO(), oidcCode)
		if err != nil {
			errMsg := fmt.Sprintf("error oauth2Config.Exchange: %v, oidcCode: %v", err, oidcCode)
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
		userInfo, err := ParseIDToken(idToken)
		log.Printf("user email: %v", userInfo.Email)

		accessToken, err := verifier.Verify(context.TODO(), oauth2Token.AccessToken)

		resp := struct {
			OAuth2Token *oauth2.Token
			AccessToken any
			IDToken     any
		}{
			OAuth2Token: oauth2Token,
			AccessToken: ParseFullIDToken(accessToken),
			IDToken:     ParseFullIDToken(idToken),
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

func randomString(len int) string {
	b := make([]byte, len)
	_, _ = rand.Read(b)
	r := hex.EncodeToString(b)
	return r
}

const StateCookieName = "state_daominah"

type MyUserInfo struct {
	KeycloakUserID string `json:"sub"`
	Email          string `json:"email"`
	EmailVerified  string `json:"email_verified"`
}

func ParseIDToken(idToken *oidc.IDToken) (MyUserInfo, error) {
	var u MyUserInfo
	err := idToken.Claims(&u)
	if err != nil {
		return u, fmt.Errorf("idToken.Claims: %v", err)
	}
	return u, nil
}

func ParseFullIDToken(idToken *oidc.IDToken) map[string]any {
	var u map[string]any
	_ = idToken.Claims(&u)
	return u
}

func AuthImplicitFlowURL(c oauth2.Config, state string) string {
	var buf bytes.Buffer
	buf.WriteString(c.Endpoint.AuthURL)
	v := url.Values{"response_type": {"token"}, "client_id": {c.ClientID}}
	if c.RedirectURL != "" {
		v.Set("redirect_uri", c.RedirectURL)
	}
	if len(c.Scopes) > 0 {
		v.Set("scope", strings.Join(c.Scopes, " "))
	}
	if state != "" {
		v.Set("state", state)
	}
	if strings.Contains(c.Endpoint.AuthURL, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	return buf.String()
}
