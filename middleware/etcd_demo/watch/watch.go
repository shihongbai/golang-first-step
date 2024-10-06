package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect etcd failed, err:%v\n", err)
		return
	}

	defer cli.Close()

	// watch
	watch := cli.Watch(context.Background(), "test")

	for watChan := range watch {
		for _, ev := range watChan.Events {
			fmt.Printf("type: %s key %s value %s", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}

}
