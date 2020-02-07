package main

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Job() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	url0 := "https://openapi.kbsec.com.vn/sso/oauth/token"
	r, err := http.NewRequest("POST", url0, nil)
	if err != nil {
		log.Println(err)
		return
	}
	w, err := httpClient.Do(r)
	if err != nil {
		log.Println(err)
		return
	}
	bodyB, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Println(err)
		return
	}
	w.Body.Close()
	log.Printf("status: %v, body: %v\n", w.Status, string(bodyB))
}

func main() {
	ticker := time.NewTicker(1 * time.Second)
	for {
		go Job()
		<-ticker.C
	}
}
