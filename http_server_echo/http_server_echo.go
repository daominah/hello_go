package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/", func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Printf("_____________________________________________\n")
			log.Printf("begin HandleFunc %v\n", r.RemoteAddr)
			rDump, err := httputil.DumpRequest(r, true)
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(fmt.Sprintf("error DumpRequest: %v", err)))
				return
			}
			w.Write([]byte("Echo of your request:\n\n"))
			w.Write(rDump)
			log.Printf("respond to %v: %s\n", r.RemoteAddr, rDump)
			log.Printf("end HandleFunc %v\n", r.RemoteAddr)
			log.Printf("_____________________________________________\n")
			return
		}
	}())

	server := &http.Server{Addr: ":20891", Handler: handler}
	log.Println("listening on port ", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	// curl 127.0.0.1:20891 --request POST --data '{"name":"Tung"}'
	// 127.0.0.1:20891?a=5

	select {}
}
