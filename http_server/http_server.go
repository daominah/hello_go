package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)

	handler := http.NewServeMux()
	handler.HandleFunc("/long-task",
		func(w http.ResponseWriter, r *http.Request) {
			go func() {
				reqDeadline, ok := r.Context().Deadline()
				log.Printf("req ctx deadline %v, %v\n", reqDeadline, ok)
				<-r.Context().Done()
				log.Println("client closed or server responded")
			}()
			log.Println("handling long-task")
			time.Sleep(2 * time.Second)
			n, err := w.Write([]byte("my slow response"))
			if err != nil {
				log.Println("error when ResponseWriter Write: ", err)
			}
			log.Printf("responded to long-task. n: %v\n", n)
		})
	server := &http.Server{Addr: ":8008", Handler: handler}

	log.Println("listening on port ", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
