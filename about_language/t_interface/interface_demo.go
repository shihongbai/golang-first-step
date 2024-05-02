package t_interface

import "fmt"

type tester interface {
	test()
	string() string
}

type data struct {
}

func (*data) test() {

}

func (data) string() string {
	return ""
}

func main() {
	var d data

	// 由于 test() 方法是由 *data 实现的，所以 data 类型的变量 d 无法直接赋值给接口变量 t，因为 d 不是指针类型
	//var t tester = d

	var t tester = &d // 将 d 的地址赋值给接口变量 t
	t.test()
	fmt.Println(t.string())
}
