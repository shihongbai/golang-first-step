package stack_array

import (
	"golang-first-step/about_leetcode/model"
	"golang.org/x/exp/constraints"
)

// 两个栈实现队列

type TowStackQueue[T constraints.Ordered] struct {
	stackPush *model.Stack[T] // 入队列的栈
	stackPop  *model.Stack[T] // 出队列的栈
}

func (t *TowStackQueue[T]) pushToPop() {
	if t.stackPop.IsEmpty() {
		for !t.stackPush.IsEmpty() {
			v, _ := t.stackPop.Pop()
			t.stackPop.Push(v)
		}
	}
}

func (t *TowStackQueue[T]) Peek() (T, bool) {
	return t.stackPop.Peek()
}

// Poll 出队列
func (t *TowStackQueue[T]) Poll() (T, bool) {
	if t.stackPop.IsEmpty() && t.stackPush.IsEmpty() {
		var zeroValue T
		return zeroValue, false
	}

	t.pushToPop()
	return t.stackPop.Pop()
}

func (t *TowStackQueue[T]) Add(e T) {
	t.stackPush.Push(e)
	t.pushToPop()
}

func NewTowStackQueue[T constraints.Ordered]() *TowStackQueue[T] {
	return &TowStackQueue[T]{}
}
