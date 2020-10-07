package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		cmd := exec.Command("/bin/bash", "-c", `free -m | awk 'NR==2{printf "Memory Usage: %s/%sMB (%.2f%%)\n", $3,$2,$3*100/$2 }'
df -h | awk '$NF=="/"{printf "Disk Usage: %d/%dGB (%s)\n", $3,$2,$5}'
grep 'cpu ' /proc/stat | awk '{usage=($2+$4)*100/($2+$4+$5)} END {print "CPU Usage: " usage "%"}'`)
		stdout, err := cmd.Output()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf("error exec: %v", err)))
			return
		}
		w.Write(stdout)
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
