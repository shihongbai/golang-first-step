package variable

import "fmt"

func slice_array() {
	// 声明一个切片
	x := make([]int, 5)
	fmt.Println("x: ", x)

	// 声明一个定长数组
	var y = [5]int{1, 2, 4, 5, 6}
	fmt.Println("y: ", y)
}
