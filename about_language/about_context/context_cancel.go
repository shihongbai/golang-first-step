package about_context

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 演示Context的全局取消信号的作用

var wg sync.WaitGroup

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second * 1)
		select {
		// 等待上级的取消信号
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
}

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 3)
	cancelFunc()
	wg.Wait()
	fmt.Println("over")
}
