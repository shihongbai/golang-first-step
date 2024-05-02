package main

import "fmt"

type N int

func (n N) test() {
	fmt.Printf("test.n %p, %d\n", &n, n)
}

func method_expression() {
	// 方法表达式的使用
	var n N = 25
	fmt.Printf("main.n: %p, %d\n", &n, n)

	f1 := N.test
	f1(n)

	// 方法表达式是一种构造函数值的语法，它包含了接收者类型，这里的 (*N) 表示指针类型 N。
	// 这个表达式构造了一个函数，它需要一个 *N 类型的接收者，并调用 test 方法
	// 所以，f2 := (*N).test 和 f2(&n) 的组合，实际上是手动指定了接收者类型为 *N，然后调用了 test 方法
	f2 := (*N).test
	f2(&n)
}

func main() {
	var n1 N = 100
	p := &n1

	// f1() 被调用时，因为 f1 是方法值，它关联的接收者是 n1 的一个副本，而不是原始的 n1。所以它打印的值仍然是 100
	n1++
	f1 := n1.test

	// 复制*p，等于102
	n1++
	f2 := p.test

	n1++
	fmt.Printf("main.n: %p, %d\n", &n1, n1)

	f1()
	f2()
}
