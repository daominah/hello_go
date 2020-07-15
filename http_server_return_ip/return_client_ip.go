package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/", func() http.HandlerFunc {
		myHostName, _ := os.Hostname()
		myLocalIP, _ := GetOutboundIP()
		log.Printf("server info: %v, %v", myHostName, myLocalIP)
		return func(w http.ResponseWriter, r *http.Request) {
			jsonB, _ := json.MarshalIndent(map[string]interface{}{
				"ServerHostname": myHostName,
				"ServerLocalIP":  myLocalIP,
				"ClientAddr":     r.RemoteAddr,
			}, "", "\t")
			w.Write(jsonB)
		}
	}())

	server := &http.Server{Addr: ":20891", Handler: handler}
	log.Println("listening on port ", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	select {}
}

func GetOutboundIP() (string, error) {
	//  any port is ok, target does not need be real
	conn, err := net.Dial("udp", "8.8.8.8:11992")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}
