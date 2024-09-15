package stack_array

import "golang-first-step/about_leetcode/model"

// 一个栈实现另一个栈的排序
// 算法实现
// 1. 记排序栈为stack，辅助栈是help
// 2. 不断将stack中的元素弹出，弹出的元素记为cur，与help栈顶进行比较
// 3. 如果cur小于等于栈顶元素，则将cur直接压入help
// 4. 如果大于help栈顶元素，则将help元素逐一弹出，逐一压入stack，直到cur小于或等于help栈顶元素，再将cur压入help
func sortStackByStack(stack *model.Stack[int]) {

}
