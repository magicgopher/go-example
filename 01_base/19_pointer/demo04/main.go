package main

import "fmt"

func main() {
	/*
		数组指针：数组指针是指一个指向整个数组的指针。它的类型是 *[]T 或 *[<size>]T，其中 T 是数组元素类型，<size> 是数组的固定长度。
			格式：*[长度]类型

		指针数组：首先是一个数组，存储的数据类型是指针。
			格式：[长度]*类型

		示例：
			*[5]float64,指针，一个存储了5个浮点类型数据的数组的指针
			*[3]string，指针，数组的指针，存储了3个字符串
			[3]*string，数组，存储了3个字符串的指针地址的数组
			[5]*float64，数组，存储了5个浮点数据的地址的数组
			*[5]*float64，指针，一个数组的指针，存储了5个float类型的数据的指针地址的数组的指针
			*[3]*string，指针，存储了3个字符串的指针地址的数组的指针
			**[4]string，指针，存储了4个字符串数据的数组的指针的指针
			**[4]*string，指针，存储了4个字符串的指针地址的数组，的指针的指针
	*/

	// 定义一个数组
	arr1 := [3]int{1, 2, 3}
	fmt.Println(arr1)

	// 创建一个指针，用于存储数组的地址 --> 数组指针
	p1 := new([3]int) // 这里使用new函数创建指针
	fmt.Printf("类型:%T\n", p1)
	p1 = &arr1
	fmt.Printf("%p\n", p1)  // arr1数组的地址
	fmt.Printf("%p\n", &p1) // p1指针自身的地址

	// 根据数组指针操作数组
	(*p1)[0] = 100
	fmt.Println(arr1) // [100 2 3]

	p1[0] = 10        // 简化写法
	fmt.Println(arr1) // [10 2 3]

	fmt.Println("===============")

	// 创建一个指针数组
	a, b, c := 77, 88, 99
	arr2 := [3]int{a, b, c}
	arr3 := [3]*int{&a, &b, &c} // 指针数组
	fmt.Println(arr2)           // [77 88 99]
	fmt.Println(arr3)           // [0x1400009a028 0x1400009a040 0x1400009a048]
	arr2[0] = 777
	fmt.Println(arr2) // [777 88 99]
	fmt.Println(a)    // 77
	*arr3[0] = 66     // 这里操作的是指针变量指向的地址
	fmt.Println(arr3) // [0x1400000e0a8 0x1400000e0c0 0x1400000e0c8]
	fmt.Println(arr2) // [777 88 99]
	fmt.Println(a)    // 66

	b = 1111
	fmt.Println(arr2)
	fmt.Println(arr3)

	// 遍历指针数组存储的地址指向的地址所存储的值
	for i := 0; i < len(arr3); i++ {
		fmt.Println(*arr3[i])
	}
}
