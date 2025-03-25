package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	logFile *os.File
	mu      sync.Mutex
)

func main() {
	var err error
	logFile, err = os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer logFile.Close()

	handler := http.NewServeMux()

	// example request:
	// http://127.0.0.1:20991/log?odds=0.2333&nStarters=3&deckSize=40&handSize=5&min=1&max=3
	handler.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		logEntry := fmt.Sprintf("%-30v %-30v %v\n",
			time.Now().UTC().Format(time.RFC3339),
			r.RemoteAddr,
			r.URL.RawQuery)

		mu.Lock()
		defer mu.Unlock()
		if _, err := logFile.WriteString(logEntry); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	// example request:
	// http://127.0.0.1:20991/view?limit=100  // limit count from the end of the file
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
