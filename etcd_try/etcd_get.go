package main

import (
	"context"
	"log"
	"time"

	"go.etcd.io/etcd/v3/clientv3"
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
	ret, err := etcdCli.Get(ctx, "/z")
	//ret, err := etcdCli.Get(ctx, "/a", clientv3.WithPrefix())
	cxl()
	if err != nil {
		log.Fatalf("error etcdCli Get: %v", err)
	}
	log.Printf("ret: %#v", ret)
	for _, kv := range ret.Kvs {
		log.Printf("k: %s, v: %s", kv.Key, kv.Value)
	}
}
