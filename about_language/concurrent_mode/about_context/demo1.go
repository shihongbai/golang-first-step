package about_context

import (
	"context"
	"fmt"
	"time"
)

// 典型的使用cancel context的例子用于实现子协程的控制

func HandelRequest(ctx context.Context) {
	go WriteRedis(ctx)
	go WriteDatabase(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done")
			return
		default:
			fmt.Println("HandelRequest Running")
			time.Sleep(1 * time.Second)
		}
	}
}

func WriteDatabase(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteDatabase Done")
			return
		default:
			fmt.Println("WriteDatabase Running")
			time.Sleep(1 * time.Second)
		}
	}
}

func WriteRedis(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteRedis Done")
			return
		default:
			fmt.Println("WriteRedis Running")
			time.Sleep(1 * time.Second)
		}
	}
}

func test() {
	ctx, cancel := context.WithCancel(context.Background())
	go HandelRequest(ctx)

	time.Sleep(5 * time.Second)
	fmt.Println("It's time to stop all sub goroutines")
	cancel()

	time.Sleep(5 * time.Second)
}
