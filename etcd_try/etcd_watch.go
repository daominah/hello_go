package main

import (
	"context"
	"log"
	"time"

	"go.etcd.io/etcd/v3/clientv3"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)
	etcdCli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{
			"192.168.99.100:2379",
			"192.168.99.101:2379",
			"192.168.99.102:2379",
		},
		DialTimeout: 30 * time.Second,
	})
	if err != nil {
		log.Fatalf("toang: %v", err)
	}
	defer etcdCli.Close()

	ctx, cxl := context.WithCancel(context.Background())
	ctx = clientv3.WithRequireLeader(ctx)
	watchChan := etcdCli.Watch(ctx, "/new_job")
	var resp clientv3.WatchResponse
	for wresp := range watchChan {
		for _, ev := range wresp.Events {
			log.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
	cxl()
	log.Printf("watchChan closed: %v, %v", resp.Err(), ctx.Err())
}
