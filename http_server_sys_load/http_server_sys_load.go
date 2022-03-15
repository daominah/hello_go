package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	var cacheSystemLoad string
	mu := &sync.Mutex{}
	go func() {
		for {
			tmp := fmt.Sprintf("%v\n%v\n%v of %v",
				GetDiskUsage(),
				GetMemoryUsage(),
				GetCPUAverageUsage(), GetCPUModel())
			mu.Lock()
			cacheSystemLoad = tmp
			mu.Unlock()
			time.Sleep(1 * time.Second)
		}
	}()

	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		mu.Lock()
		response := cacheSystemLoad
		mu.Unlock()
		w.Write([]byte(response))
	})

	listen := os.Getenv("LISTENING_PORT")
	if listen == "" {
		listen = ":21864"
	}
	if !strings.Contains(listen, ":") {
		listen = ":" + listen
	}

	server := &http.Server{Addr: listen, Handler: handler}
	log.Printf("listening on http://127.0.0.1%v/\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
