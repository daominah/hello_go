package main

import (
	"bufio"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

func main0() {
	// web view circuit breaker metric in rolling window (10 seconds)
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(":20211", hystrixStreamHandler)

	// config the circuit breaker
	command0 := "ReadHistory"
	hystrix.ConfigureCommand(command0, hystrix.CommandConfig{
		Timeout:                1300, // milliseconds
		MaxConcurrentRequests:  20,
		RequestVolumeThreshold: 10,    // minimum nRequests to calc ErrorPercent
		SleepWindow:            30000, // milliseconds to wait after circuit open
		ErrorPercentThreshold:  50,
	})

	// try to query database
	db := &DatabaseMock{}
	rand.Seed(time.Now().UnixNano())

	wg := &sync.WaitGroup{}
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			defer wg.Add(-1)
			run := func() error {
				rows, err := db.ReadHistory()
				_ = rows
				return err
			}
			fallback := func(hystrixErr error) error {
				switch hystrixErr {
				case hystrix.ErrMaxConcurrency:
					fallthrough
				case hystrix.ErrTimeout:
					fallthrough
				case hystrix.ErrCircuitOpen:
					fallthrough
				default:
					log.Printf("fallback: %v", hystrixErr)
				}
				return nil
			}
			hystrix.Do(command0, run, fallback)
		}()
	}
	wg.Wait()

	// view metric
	w, _ := http.Get("http://127.0.0.1:20891/")
	reader := bufio.NewReader(w.Body)
	metric, _ := reader.ReadBytes('\n')
	w.Body.Close()
	log.Printf("metric: %s\n", metric)

	log.Println("done")
}

type Database interface {
	ReadHistory() ([]string, error)
}

type DatabaseMock struct{}

func (d DatabaseMock) ReadHistory() ([]string, error) {
	exeDur := time.Duration(500+rand.Intn(1000)) * time.Millisecond
	log.Printf("exeDur: %v\n", exeDur)
	time.Sleep(exeDur)
	return []string{"row2", "row1", "row0"}, nil
}

type DatabaseMock2 struct{}

func (d DatabaseMock2) ReadHistory() ([]string, error) {
	w, err := http.Get("https://google.com.vn")
	if err != nil {
		return nil, err
	}
	w.Body.Close()
	return []string{"row2", "row1", "row0"}, nil
}
