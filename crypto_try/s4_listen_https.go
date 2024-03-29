package main

import (
	"log"
	"net/http"
)

func _main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("request from %v", r.RemoteAddr)
			w.Write([]byte("ngon"))
		})
	s := http.Server{Addr: ":43443", Handler: handler}
	log.Print("listening on https://127.0.0.1:43443/")
	err := s.ListenAndServeTLS(
		"/home/tungdt/go/src/github.com/daominah/hello_go/crypto_try/myorg0.crt",
		"/home/tungdt/go/src/github.com/daominah/hello_go/crypto_try/myorg0.key")
	if err != nil {
		log.Fatal(err)
	}
}
