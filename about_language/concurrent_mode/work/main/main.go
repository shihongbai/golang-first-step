// work包的使用案例
package main

import (
	"golang-first-step/about_language/concurrent_mode/work"
	"log"
	"sync"
	"time"
)

// names 创建一组用于显示的名字
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

type namePrinter struct {
	name string
}

// Task 实现Worker接口
func (n *namePrinter) Task() {
	log.Print(n.name)
	time.Sleep(time.Second)
}

func main() {
	// 使用两个goroutine来创建工作池
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(10 * len(names))
	for i := 0; i < 100; i++ {
		// 迭代names切片
		for _, name := range names {
			// 创建一个namePrinter并提供指定名字
			np := namePrinter{
				name: name,
			}

			go func() {
				// 将任务提交执行，当Run返回时候
				// 就确定任务已经处理完成
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	// 让工作池停止工作，等待所有现有的工作完成
	p.Shutdown()
}
