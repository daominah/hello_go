package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/mywrap/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

type Row struct {
	Key string `bson:"key"`
	Val string `bson:"val"`
}

func main() {
	dbName := "cafe"
	collHundred := "hundred"
	collBillion := "billion"
	_, _ = collBillion, collHundred
	conf := mongodb.Config{Host: "127.0.0.1", Port: "27017", Database: dbName}
	dbPool, err := mongodb.Connect(conf)
	if err != nil {
		log.Fatalf("error when open mongodb db: %v", err)
	}

	if false { // create big collection, can take several hours
		n := 1000000000
		nBulk := 100
		bulkSize := n / nBulk
		for bi := 0; bi < nBulk; bi++ {
			log.Printf("bulk i: %v\n", bi)
			rows := make([]interface{}, 0)
			for i := bi * bulkSize; i < (bi+1)*bulkSize; i++ {
				rows = append(rows, Row{
					Key: fmt.Sprintf("%v", i),
					Val: fmt.Sprintf("%v", i),
				})
			}
			log.Println("about to shuffle")
			rand.Shuffle(bulkSize, func(i, j int) { rows[i], rows[j] = rows[j], rows[i] })

			log.Println("about to insert")
			if bi == 0 {
				_, err = dbPool.Database(dbName).Collection(collHundred).InsertMany(
					context.Background(), rows[:100])
				if err != nil {
					log.Fatal(err)
				}
			}

			_, err = dbPool.Database(dbName).Collection(collBillion).InsertMany(
				context.Background(), rows)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if true { // try insert and select 100 times from big collection
		sumDur := time.Duration(0)
		for l := 200; l < 300; l++ {
			beginInsertT := time.Now()
			i := 4000000000 + l
			key := fmt.Sprintf("%v", i)
			_, err = dbPool.Database(dbName).Collection(collBillion).InsertOne(
				context.Background(), Row{Key: key, Val: key})
			log.Println("insert result:", err)
			inDur := time.Since(beginInsertT)
			log.Println("insert dur:", inDur)
			beginSelectT := time.Now()
			cursor, err := dbPool.Database(dbName).Collection(collBillion).Find(
				context.Background(), bson.M{"key": key})
			log.Println("select result:", err)
			for cursor.Next(context.Background()) {
				log.Println("a row:", cursor.Current.String())
			}
			seDur := time.Since(beginSelectT)
			log.Println("select dur:", seDur)
			sumDur += inDur + seDur
		}
		log.Println("sumDur", sumDur)
	}
}
