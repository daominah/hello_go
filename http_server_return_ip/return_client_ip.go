package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			jsonB, _ := json.Marshal(map[string]interface{}{
				"ClientAddr": fmt.Sprintf("%v", r.RemoteAddr),
			})
			w.Write(jsonB)
		})
	server := &http.Server{Addr: ":33066", Handler: handler}
	log.Println("listening on port ", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	select {}
}
