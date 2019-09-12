package main

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"time"
)

func main() {

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
	})
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := "tuxedo"
	for i := 0; i > -1; i++ {
		time.Sleep(1 * time.Millisecond)
		for _, word := range []string{`{"messageType":"REQUEST","sourceId":"fix","transactionId":null,"messageId":"2c546836-9121-430b-8ec9-2f28827554ca","uri":"/api/v1/equity/order/history","responseDestination":{"topic":"fix.5fc72175-5727-4fe1-bab5-2f5b8a6ed688","uri":"/api/v1/equity/order/history"},"data":{"headers":{"token":{"domain":"kis","userId":"","serviceCode":"kis","connectionId":null,"serviceId":null,"serviceName":null,"clientId":null,"serviceUserId":4,"loginMethod":20,"refreshTokenId":null,"scopeGroupIds":[1,8,10,12,15],"serviceUsername":"057ECB5693","userData":{"username":"057ECB5693","identifierNumber":"56/UBCK_GPH?KD","branchCode":"00","mngDeptCode":"100","deptCode":"100","agencyNumber":"000","caThumbprint":null,"accountNumbers":["057ECB5693"]}},"secToken":null,"acceptLanguage":null},"sourceIp":null,"deviceType":null,"accountNumber":"057ECB5693","subNumber":"00","fromDate":"20190912","toDate":null,"stockCode":null,"sellBuyType":null,"matchType":null,"sortType":"ASC","lastOrderDate":"20190912","lastBranchCode":"100","lastOrderNumber":"1358","lastMatchPrice":0.0,"fetchCount":100,"marketType":null},"t":null,"stream":null,"streamState":null,"streamIndex":null}`} {
			err := producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{
					Topic:     &topic,
					Partition: kafka.PartitionAny},
				Value: []byte(fmt.Sprintf("%v", word)),
			}, nil)
			if err != nil {

			}
		}
	}

	// Wait for message deliveries before shutting down
	producer.Flush(1000 * 1000)
}
