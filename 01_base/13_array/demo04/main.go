package main

import "fmt"

func main() {
	// 定义一个string类型的数组
	var arr1 [5]string
	// 获取数组长度
	fmt.Printf("arr1数组的长度:%d\n", len(arr1))
	fmt.Printf("arr1数组的容量:%d\n", cap(arr1))
	// 通过数组索引下标访问数组元素
	arr1[0] = "Go"
	arr1[1] = "Java"
	arr1[2] = "Python"
	arr1[3] = "C++"
	arr1[4] = "Rust"
	// 通过for循环遍历数组的元素
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("arr1[%d]=%s\n", i, arr1[i])
	}
}
