package model

import "golang.org/x/exp/constraints"

// Stack 定义泛型栈结构体，要求类型T必须是可比较大小的
type Stack[T constraints.Ordered] struct {
	elements []T
}

// NewStack 构造方法，返回一个新的栈实例
func NewStack[T constraints.Ordered]() *Stack[T] {
	return &Stack[T]{
		elements: make([]T, 0), // 初始化空切片
	}
}

// Push 添加元素到栈顶
func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

// Pop 从栈顶弹出元素
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	value := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return value, true
}

// Peek 查看栈顶元素但不弹出
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.elements) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	return s.elements[len(s.elements)-1], true
}

// FindMax 找到栈中最大的元素
func (s *Stack[T]) FindMax() (T, bool) {
	if len(s.elements) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	max := s.elements[0]
	for _, v := range s.elements {
		if v > max {
			max = v
		}
	}
	return max, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}
