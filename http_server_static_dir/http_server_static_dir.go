package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// export GO111MODULE=off && /usr/local/go/bin/go build
func main() {
	SetLogTimeISO()

	var dir, port string
	flag.StringVar(&dir, "i", "", "The directory to serve static files from")
	flag.StringVar(&port, "p", "", "The port to run the server on")
	flag.Parse()
	if dir == "" {
		dir = "." // carefull, this can be security issue, e.g. unintended serve .git folder
		log.Printf("empty flag i (dir to serve), use current dir")
	} else {
		log.Printf("flag i (dir to serve): %v", dir)
	}
	port = strings.TrimPrefix(port, ":")
	if port == "" {
		port = "28384"
		log.Printf("empty flag p (port), use default port %v", port)
	} else {
		log.Printf("flag p (port): %v", port)
	}

	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)
	log.Printf("serving %v on http://localhost:%v", dir, port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

type LogWriterISOTime struct{}

func (w LogWriterISOTime) Write(output []byte) (int, error) {
	return fmt.Fprintf(os.Stderr, "%v %s", time.Now().UTC().Format(time.RFC3339), output)
}

func SetLogTimeISO() {
	log.SetFlags(log.Lshortfile)
	log.SetOutput(LogWriterISOTime{})
}
