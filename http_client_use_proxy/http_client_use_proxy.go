package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	proxyUrl, err := url.Parse("http://27.147.209.215:8080")
	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy:           http.ProxyURL(proxyUrl),
			MaxConnsPerHost: 6,
		},
		Timeout: 20 * time.Second,
	}

	r, _ := http.NewRequest("GET", "http://icanhazip.com/", nil)
	w, err := httpClient.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer w.Body.Close()
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}
	_ = body
	log.Printf("body: %s", body)
}
