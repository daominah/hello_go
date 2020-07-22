package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
	server := &http.Server{Addr: ":20891"}
	go server.ListenAndServe()

	var httpClient *http.Client
	if false { // create max 100 conns, reuse for next requests
		httpClient = &http.Client{
			Transport: &http.Transport{MaxConnsPerHost: 100},
			Timeout:   10 * time.Second,
		}
	} else { // watch -n 0.2 "netstat | wc -l"
		httpClient = http.DefaultClient
	}

	wg := &sync.WaitGroup{}
	for k := 0; k < 100000; k++ {
		wg.Add(1)
		go func() {
			defer wg.Add(-1)
			sendRequest(httpClient)
		}()
	}
	wg.Wait()
	log.Println("done")
}

func sendRequest(httpClient *http.Client) error {
	r, _ := http.NewRequest("GET", "http://127.0.0.1:20891/", nil)
	w, err := httpClient.Do(r)
	if err != nil {
		log.Println(err)
		return err
	}
	defer w.Body.Close()
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		return err
	}
	_ = body
	return nil
}
