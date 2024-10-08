package main

import (
	"fmt"
	"time"
)

// select中各个case执行顺序是随机的，如果某个case中的channel已经ready，则执行相应的语句并退
// 出select流程，如果所有case中的channel都未ready，则执行default中的语句然后退出select流程。另外，
// 由于启动的协程和select语句并不能保证执行顺序，所以也有可能select执行时协程还未向channel中写入数据，
// 所以select直接执行default语句并退出。因此最终的执行结果有三种
func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		chan1 <- 1
		time.Sleep(1 * time.Second)
	}()

	go func() {
		chan2 <- 1
		time.Sleep(1 * time.Second)
	}()

	select {
	case <-chan1:
		fmt.Println("chain1 ready")
	case <-chan2:
		fmt.Println("chain2 ready")
	default:
		fmt.Println("default")
	}

	fmt.Println("main exit")
}
