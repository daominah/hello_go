package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/mywrap/log"
	"github.com/mywrap/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const cbCreateRow = "createRow"

func init() {
	hystrix.ConfigureCommand(cbCreateRow, hystrix.CommandConfig{
		Timeout: 600, // milliseconds
		//Timeout: 1990, // percentile 99.5, will not open the circuit

		MaxConcurrentRequests:  15,
		RequestVolumeThreshold: 10,    // minimum nRequests to calc ErrorPercent
		SleepWindow:            10000, // milliseconds to wait after circuit open
		ErrorPercentThreshold:  25,
	})
	// web view circuit breaker metric in rolling window (10 seconds)
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(":20211", hystrixStreamHandler)
}

type Repo struct {
	DB *mongo.Client
}

func (r Repo) myColl() *mongo.Collection {
	return r.DB.Database("database0").Collection("collection0")
}

func (r Repo) createRow(isFast bool) (string, error) {
	var retId string
	var retErr error
	callDep := func() error {
		log.Printf("about to call database")
		rowData := bson.M{"lastModified": time.Now()}
		if rand.Intn(100) < 10 {
			rowData = bson.M{"lastModified": func() {}} // invalid data
		}
		qr, err := r.myColl().InsertOne(context.Background(), rowData)
		if !isFast {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		}
		if err != nil {
			retErr = err
			return nil
		}
		retId = fmt.Sprintf("%v", qr.InsertedID)
		return nil
	}
	cbErr := hystrix.Do(cbCreateRow, callDep, nil)
	if cbErr == nil {
		return retId, retErr
	}
	// fallback
	return "cacheData", cbErr
}

func main() {
	client, err := mongodb.Connect(mongodb.LoadEnvConfig())
	if err != nil {
		log.Fatal(err)
	}
	r := Repo{DB: client}
	for i := 0; true; i++ {
		time.Sleep(30 * time.Millisecond)
		if i%100 == 99 {
			log.Printf("i: %v", i)
		}
		if i%500 == 499 {
			time.Sleep(15000 * time.Millisecond)
		}
		go func(isFast bool) {
			retId, err := r.createRow(isFast)
			if err != nil {
				log.Printf("error createRow: %v, retId: %v", err, retId)
				return
			} else {

			}
		}(i%500 < 10)
	}
	log.Printf("done")
}
