package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"h12.io/socks"
)

func main1() { // standard lib, does not support socks4
	// https://gimmeproxy.com/api/getProxy?protocol=socks5
	proxyUrl, err := url.Parse("socks5://20.184.5.198:1080")
	cli := &http.Client{
		Transport: &http.Transport{
			// "http", "https", and "socks5" are supported
			Proxy:           http.ProxyURL(proxyUrl),
			MaxConnsPerHost: 6,
		},
		Timeout: 20 * time.Second,
	}

	r, _ := http.NewRequest("GET", "http://icanhazip.com/", nil)
	status, body, err := httpDo(cli, r)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("status: %v, body:\n%s\n", status, body)
}

func main() {
	dialFunc := socks.Dial("socks4://36.92.9.75:49420?timeout=15s")
	cli := &http.Client{Transport: &http.Transport{Dial: dialFunc}}

	r, _ := http.NewRequest("GET", "http://136.144.56.255", nil)
	status, body, err := httpDo(cli, r)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("status: %v, body:\n%s\n", status, body)
}

func httpDo(cli *http.Client, r *http.Request) (
	respStatus string, respBody []byte, retErr error) {
	w, err := cli.Do(r)
	if err != nil {
		retErr = fmt.Errorf("error client Do: %v", err)
		return
	}
	defer w.Body.Close()
	respStatus = fmt.Sprintf("%v: %v", w.StatusCode, w.Status)
	respBytes, err := ioutil.ReadAll(w.Body)
	if err != nil {
		retErr = fmt.Errorf("error ioutil ReadAll: %v", err)
		return
	}
	respBody = respBytes
	return
}
