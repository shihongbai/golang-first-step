package _range

// 数通过遍历切片，打印切片的下标和元素值，请问性能上有没有可优化的空间？
func rangeSlice(arr []int) {
	//遍历过程中每次迭代会对index和value进行赋值，如果数据量大或者value类型为string时，对value
	//的赋值操作可能是多余的，可以在for-range中忽略value值，使用slice[index]引用value值
	for index, value := range arr {
		_, _ = index, value
	}
}

// 函数通过遍历Map，打印Map的key和value，请问性能上有没有可优化的空间？
func rangeMap(myMap map[int]string) {
	//函数中for-range语句中只获取key值，然后跟据key值获取value值，虽然看似减少了一次赋值，但通
	//过key值查找value值的性能消耗可能高于赋值消耗。能否优化取决于map所存储数据结构特征、结合实际情况进行
	for key, _ := range myMap {
		_, _ = key, myMap[key]
	}
}

// 请问如下程序是否能正常结束？
func dynamicRange() {
	v := []int{1, 2, 3, 4, 5}
	// 能够正常结束。循环内改变切片的长度，不影响循环次数，循环次效在循环开始前就已经确定了
	for i := range v {
		v = append(v, i)
	}
}
