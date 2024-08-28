package stack_array

import (
	"golang-first-step/about_leetcode/model"
	"golang.org/x/exp/constraints"
)

// 设计一个具有getMin功能的栈
// 设计要求
// 1. pop、push、getMin操作都是O(1)
// 2. 设计栈的栈类型可以使用现成的栈结构

type MinStack[T constraints.Ordered] struct {
	stackData *model.Stack[T] // 实际数据存储栈
	stackMin  *model.Stack[T] // 最小栈
}

func (m *MinStack[T]) GetMin() (T, bool) {
	return m.stackMin.Peek()
}

func (m *MinStack[T]) Push(e T) {
	if m.stackData.IsEmpty() {
		m.stackData.Push(e)
		m.stackMin.Push(e)
		return
	}

	// 检查stackMin栈顶元素
	peek, _ := m.stackMin.Peek()
	if peek <= e {
		m.stackData.Push(e)
		m.stackMin.Push(peek)
	} else {
		m.stackData.Push(e)
		m.stackMin.Push(e)
	}
}

func (m *MinStack[T]) Pop() (T, bool) {
	// 检查是否栈空
	if m.stackData.IsEmpty() {
		return m.stackData.Pop()
	}

	m.stackMin.Pop()
	return m.stackData.Pop()
}

func NewMinStack[T constraints.Ordered]() *MinStack[T] {
	return &MinStack[T]{
		stackData: model.NewStack[T](),
		stackMin:  model.NewStack[T](),
	}
}
