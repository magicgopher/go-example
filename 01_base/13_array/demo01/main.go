package main

import "fmt"

func main() {
	// 数组：一组固定长度且存储同一种数据类型的数据结构

	// 定义数组
	var arr1 [5]int
	fmt.Printf("%T, %v\n", arr1, arr1)

	// 定义数组并初始化
	arr2 := [3]string{"Go", "Java", "Python"}
	fmt.Printf("%T, %v\n", arr2, arr2)

	// 根据元素个数自动推导数组长度
	arr3 := [...]float64{9.99, 5.55, 8.8}
	fmt.Printf("%T, %v\n", arr3, arr3)
}
