package main

import "fmt"

func main() {
	// go的数值是值类型，存储的是值本身，将数据传递给其他的变量，传递的是数据的副本（备份）
	var num1 = 100
	fmt.Printf("类型:%T, 值是:%d\n", num1, num1)
	num2 := num1
	fmt.Printf("类型:%T, 值是:%d\n", num2, num2)
	num1 = 222 // 修改num1的值
	fmt.Printf("num1的值:%d, num2的值:%d\n", num1, num2)

	// 创建数组
	var array1 = [3]int{111, 222, 333}
	array2 := array1
	fmt.Printf("类型:%T, 地址:%p, 值:%v\n", array1, &array1, array1)
	fmt.Printf("类型:%T, 地址:%p, 值:%v\n", array2, &array2, array2)
	fmt.Printf("array1地址和array2地址是否一致，结果:%t\n", &array1 == &array2)

	array1[0] = 100 // 修改array1数组的第一个元素
	fmt.Printf("array1的值:%v\n", array1)
	fmt.Printf("array2的值:%v\n", array2)
	fmt.Printf("array1地址和array2地址是否一致，结果:%t\n", &array1 == &array2)
}
