package main

import "fmt"

func main() {
	// 创建数组
	var arr1 [3]int
	// 尝试将长度为 5 的数组赋值给长度为 3 的数组
	//arr1 = [5]int{1, 2, 3, 4, 5}
	// 给数组赋值
	arr1 = [3]int{1, 2, 3}
	fmt.Printf("%T, %v\n", arr1, arr1)
}
