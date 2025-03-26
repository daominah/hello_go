package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	logFile    *os.File
	lineNumber int
	mu         sync.Mutex
)

func main() {
	const logFileName = "server.log"

	// just open, count the number of lines, then close
	tmpF, err := os.Open(logFileName)
	if err == nil { // if the file exists
		scanner := bufio.NewScanner(tmpF)
		for scanner.Scan() {
			lineNumber++
		}
		tmpF.Close()
	}

	// reopen the file as a global var, only close when the server is stopped
	logFile, err = os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		log.Fatalf("error open log file: %v", err)
	}
	defer logFile.Close()

	tmp, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
	if err != nil {
		log.Fatalf("error init VN time zone: %v", err)
	}
	vnTimeZone := tmp.Location()

	handler := http.NewServeMux()
	allowedOrigins := []string{
		"http://localhost", // dangerous, only for testing
		"https://yugiodd.github.io",
	}

	// example request:
	// http://localhost:20991/log?odds=0.2333&nStarters=3&deckSize=40&handSize=5&min=1&max=3
	handler.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		origin := r.Header.Get("Origin")

		isAllowed := origin == ""
		for _, allowedOrigin := range allowedOrigins {
			if strings.Contains(origin, allowedOrigin) {
				isAllowed = true
				break
			}
		}
		if !isAllowed {
			http.Error(w, "Origin not allowed", http.StatusForbidden)
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET")

		remoteAddr := r.RemoteAddr
		// handle if the server is behind cloudflare tunnel or a proxy
		if forwardedFor := r.Header.Get("X-Forwarded-For"); forwardedFor != "" {
			remoteAddr = forwardedFor
		} else if cfConnectingIP := r.Header.Get("CF-Connecting-IP"); cfConnectingIP != "" {
			remoteAddr = cfConnectingIP
		}

		mu.Lock()
		lineNumber++
		logEntry := fmt.Sprintf("%-7v%-25v %-43v %v\n",
			lineNumber,
			time.Now().In(vnTimeZone).Format(time.RFC3339),
			remoteAddr,
			r.URL.RawQuery)
		if _, err := logFile.WriteString(logEntry); err != nil {
			mu.Unlock()
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		mu.Unlock()

		w.WriteHeader(http.StatusOK)
	})

	// example request:
	// http://localhost:20991/view?limit=50  // limit count from the end of the file
	handler.HandleFunc("/view", func(w http.ResponseWriter, r *http.Request) {
		limit := 100
		if l := r.URL.Query().Get("limit"); l != "" {
			if lInt, err := strconv.Atoi(l); err == nil {
				limit = lInt
			}
		}

		mu.Lock()
		defer mu.Unlock()

		file, err := os.Open("server.log")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		lines := []string{}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		start := len(lines) - limit
		if start < 0 {
			start = 0
		}

		for _, line := range lines[start:] {
			fmt.Fprintln(w, line)
		}
	})

	server := &http.Server{Addr: ":20991", Handler: handler}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server error:", err)
	}
}
