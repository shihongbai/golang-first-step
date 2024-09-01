package stack_array

import "golang-first-step/about_leetcode/model"

// 通过递归实现逆序栈
// 实现两个递归函数
// 1. 将栈底元素移除并返回
func getAndRemoveLastElement(stack model.Stack[int]) int {
	v, _ := stack.Pop()
	if stack.IsEmpty() {
		return v
	} else {
		lastV := getAndRemoveLastElement(stack)
		// 不断递归调用，使得栈底的元素返回
		stack.Push(v)
		return lastV
	}
}

// 2. 逆序一个栈
func reserve(stack model.Stack[int]) {
	if stack.IsEmpty() {
		return
	}

	// 不断返回栈底元素
	lastElement := getAndRemoveLastElement(stack)
	reserve(stack)
	// 重新反向压入栈
	stack.Push(lastElement)
}
