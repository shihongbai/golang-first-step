package pointer

import "fmt"

func allocMem(size int) *int {
	p := make([]int, size)
	return &p[0]
}

// 修改指针指向
func swap(p1, p2 *int) {
	*p1, *p2 = *p2, *p1
}

// 二级指针的使用
func updatePtrPtr(pptr **int) {
	val := 100
	*pptr = &val
}

func main() {
	var ptr *int
	updatePtrPtr(&ptr)
	fmt.Println(*ptr) // 输出：100
}
