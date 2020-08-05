package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"testing"
	"time"
)

func Test_HttpsListen(t *testing.T) {
	log.SetFlags(log.Ltime | log.Lshortfile)

	handler := http.NewServeMux()
	handler.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) { // echo
			reqDumped, _ := httputil.DumpRequest(r, true)
			resp := fmt.Sprintf("received at %v request:\n%s",
				time.Now().Format(time.RFC3339Nano), reqDumped)
			w.Write([]byte(resp))
		})

	server := &http.Server{Addr: ":44433", Handler: handler}
	go func() {
		log.Println("listening on https://127.0.0.1" + server.Addr)
		err := server.ListenAndServeTLS("./cert.pem", "./key.pem")
		if err != nil {
			log.Fatal(err)
		}
	}()
	time.Sleep(5 * time.Minute)
}
