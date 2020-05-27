package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	client := http.Client{Timeout: 1 * time.Second}
	r, _ := http.NewRequest("GET", "http://127.0.0.1:8008/long-task", nil)
	w, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer w.Body.Close()
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("response: %s\n", body)
}
