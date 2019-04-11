package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func SendRequest(httpClient *http.Client) error {
	request, _ := http.NewRequest("GET", "http://127.0.0.1:8080/", nil)
	response, err := httpClient.Do(request)
	if err != nil {
		log.Println(err)
		return err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	_ = body
	response.Body.Close()
	return nil
}

func SendAlotRequests(isGood bool) {
	_ = time.Second
	waiter := sync.WaitGroup{}
	for i := 0; i<100; i ++ {
		waiter.Add(1)
		go func () {
			var httpClient *http.Client
			if isGood {
				//this work
				httpClient = &http.Client{
					Transport: &http.Transport{MaxIdleConns: 10},
					Timeout:   10 * time.Second,
				}
			} else {
				// this does not work
				httpClient = http.DefaultClient
			}
			for i := 0; i < 10000; i++ {
				if i%1000 == 0 {
					log.Println(i)
				}
				SendRequest(httpClient)
			}
			waiter.Add(-1)
		}()
	}
	waiter.Wait()
}

func
main() {
	go http.ListenAndServe(":8080", nil)
	isGood := false
	//isGood = true
	SendAlotRequests(isGood)
}
