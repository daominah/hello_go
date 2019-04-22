package main

import (
	"fmt"
	"strings"
	"time"

	//	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

const (
	SPLITTER = "\n____________________________________________________________\n"
)

func DumpRequest(r *http.Request) string {
	fmt.Println("test %+v", r)
	part1 := fmt.Sprintf("%v received request from %v at %v: ",
		r.Host, r.RemoteAddr, time.Now().Format(time.RFC3339))
	temp, err := httputil.DumpRequest(r, true)
	if err != nil {
		return fmt.Sprintf("ERROR: %v", err)
	}
	part2 := fmt.Sprintf(string(temp))
	part3 := fmt.Sprintf(SPLITTER)
	return strings.Join([]string{part1, part2, part3}, "\n")

	//http.
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dumpedReq := DumpRequest(r)
		fmt.Print(dumpedReq)
		w.Write([]byte(dumpedReq))
	})
	go func() {
		fmt.Println("Server started")
		err := http.ListenAndServe(":8080", nil)
		//err := http.ListenAndServeTLS(":8081",
		//	"http_server/cert.pem", "http_server/key.pem", nil)
		if err != nil {
			fmt.Println("http.ListenAndServe err", err)
		}
	}()
	select {}
}
