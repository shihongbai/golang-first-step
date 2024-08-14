package main

import (
	"fmt"
)

// select会按照随机的顺序检测各case语句中channel是否ready，考虑到已关闭的channel也是可读
// 的，所以上述程序中select不会阻塞，具体执行哪个case语句具是随机的。
func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		close(chan1)
	}()

	go func() {
		close(chan2)
	}()

	select {
	case <-chan1:
		fmt.Println("chain1 ready")
	case <-chan2:
		fmt.Println("chain2 ready")
	}

	fmt.Println("main exit")
}
