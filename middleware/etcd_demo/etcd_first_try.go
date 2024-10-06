package etcd_demo

import (
	"context"
	"fmt"
	client "go.etcd.io/etcd/clientv3"
	"time"
)

// 代码连接etcd集群
func main() {

	cli, err := client.New(client.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Printf("connect etcd failed, err:%v\n", err)
		return
	}

	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	// put
	_, err = cli.Put(ctx, "test", "hello etcd")
	if err != nil {
		fmt.Printf("put failed, err:%v\n", err)
		return
	}

	cancel()

	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	gt, err := cli.Get(ctx, "test")
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}

	for _, kv := range gt.Kvs {
		fmt.Printf("key:%s, value:%s\n", kv.Key, string(kv.Value))
	}

	defer cancel()
}
