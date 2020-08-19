package main

import (
	"context"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)

	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{
			"192.168.99.100:2379",
			"192.168.99.101:2379",
			"192.168.99.102:2379",
		},
		DialTimeout: 5 * time.Second,
		Username:    "root",
		Password:    "123qwe",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	log.Printf("connected to etcd cluseter: %v", cli.Endpoints())

	// all locks created by this session have a TTL of 3s
	// (actual TTL value is at least the TTL)
	s1, err := concurrency.NewSession(cli, concurrency.WithTTL(5))
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()

	const lockName = "/myLock2"
	m1 := concurrency.NewMutex(s1, lockName)

	s2, err := concurrency.NewSession(cli, concurrency.WithTTL(10))
	if err != nil {
		log.Fatal(err)
	}
	defer s2.Close()
	m2 := concurrency.NewMutex(s2, lockName)

	log.Printf("s1 try to acquire the lock")
	if err := m1.Lock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	log.Println("acquired lock for s1")

	go func() {
		for {
			log.Printf("s2 try to acquire the lock")
			if err := m2.Lock(context.TODO()); err != nil {
				log.Printf("error s2 Lock: %v", err)
				continue
			}
			log.Println("acquired lock for s2")
			break
		}
	}()

	time.Sleep(15 * time.Second)

	// skip this to simulate crash after lock
	for true {
		log.Printf("s1 try to unlock")
		if err := m1.Unlock(context.TODO()); err != nil {
			log.Printf("error s1 Unlock: %v", err)
			continue
		}
		log.Println("released lock for s1")
		break
	}

	time.Sleep(3600 * time.Second)
}
