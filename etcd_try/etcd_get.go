package main

import (
	"context"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
)

func main() {
	etcdCli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{
			"192.168.99.100:2379",
			"192.168.99.101:2379",
			"192.168.99.102:2379",
		},
		DialTimeout: 30 * time.Second,
		//Username:    "root",
		//Password:    "123qwe",
	})
	if err != nil {
		log.Fatalf("toang: %v", err)
	}
	defer etcdCli.Close()

	ctx, cxl := context.WithTimeout(context.Background(), 5*time.Second)
	ret, err := etcdCli.Get(ctx, "key", clientv3.WithPrefix())
	cxl()
	if err != nil {
		switch err {
		case context.Canceled:
			log.Printf("canceled: %v", err)
		case context.DeadlineExceeded:
			log.Printf("timeout: %v", err)
		case rpctypes.ErrEmptyKey:
			log.Printf("client error: %v", err)
		default:
			log.Printf("bad etcd servers: %v", err)
		}
	}
	log.Printf("ret: %#v", ret)
	for _, kv := range ret.Kvs {
		log.Printf("k: %s, v: %s", kv.Key, kv.Value)
	}
}
