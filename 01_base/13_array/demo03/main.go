package main

import "fmt"

func main() {
	// 创建数组
	arr1 := [5]int{1, 2, 3, 4, 5}

	// 通过数组[索引]的方式返回数组下标存储的元素
	// 数组索引从0开始，最大索引是数组长度-1
	fmt.Printf("arr1[0]=%d\n", arr1[0])

	// 获取数组的长度
	fmt.Printf("arr1数组的长度:%d\n", len(arr1))
}
