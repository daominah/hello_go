package main

import (
	"go/build"
	"log"
	"net/http"
)

func main() {
	handler := http.NewServeMux()
	dirPath := build.Default.GOPATH + `/src/github.com/daominah/hello_go/http_server_html/gui`
	handler.Handle("/", http.FileServer(http.Dir(dirPath)))
	handler.HandleFunc("/hello",
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("request from %v\n", r.RemoteAddr)
			w.Write([]byte("Hello " + r.RemoteAddr))
		})

	s := &http.Server{Addr: ":5000", Handler: handler}
	log.Printf("listening on http://127.0.0.1%v/\n", s.Addr)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	select {}
}
