package main

import (
	"crypto/tls"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/mywrap/log"
)

func main() {
	proxies := []string{
		"http://108.165.189.7:17654",
		"http://108.165.189.55:30410",
		"http://108.165.188.182:25255",
		"http://108.165.188.248:17799",
		"http://108.165.189.91:10921",
	}
	rand.Seed(time.Now().UnixNano())

	httpClient := &http.Client{
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy: func(request *http.Request) (*url.URL, error) {
				chosen := proxies[rand.Intn(len(proxies))]
				return url.Parse(chosen)
			},
		},
	}
	r, err := http.NewRequest("GET", "https://ipv4.icanhazip.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	if true {
		w, err := httpClient.Do(r)
		if err != nil {
			log.Fatal(err)
		}
		defer w.Body.Close()
		body, err := io.ReadAll(w.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("status: %v, body: %s", w.Status, body)
	} else {
		w, err := httpClient.Transport.RoundTrip(r)
		if err != nil {
			log.Fatal(err)
		}
		defer w.Body.Close()
		body, err := io.ReadAll(w.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("status: %v, body: %s", w.Status, body)
	}
}
