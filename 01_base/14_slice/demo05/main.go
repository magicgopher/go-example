package main

import "fmt"

func main() {
	/*
		注意事项：
			1.切片并不是数组的拷贝，它只是对数组的一个引用。因此，修改切片中的元素会影响到原始数组。
			2.切片的长度可以动态变化，而数组的长度是固定的。
			3.可以使用 len() 函数获取切片的长度，使用 cap() 函数获取切片的容量。
	*/
	// 创建数组
	array := [5]int{1, 2, 3, 4, 5}

	// 将整个数组转为切片
	slice1 := array[:]
	fmt.Println("slice1:", slice1)                                 // 输出：slice1: [1 2 3 4 5]
	fmt.Printf("slice1长度: %d, 容量: %d\n", len(slice1), cap(slice1)) // 输出：slice1长度: 5, 容量: 5

	// 将数组的一部分转换为切片
	slice2 := array[1:4]
	fmt.Println("slice2:", slice2)                                 // 输出：slice2: [2 3 4]
	fmt.Printf("slice2长度: %d, 容量: %d\n", len(slice2), cap(slice2)) // 输出：slice2长度: 3, 容量: 4

	fmt.Println("==========")

	// 数组属于值类型，存储的是数据本身，传递的是数据副本
	array1 := [5]int{7, 8, 9, 10, 11}
	fmt.Printf("类型:%T, 地址:%p, 值:%v\n", array1, &array1, array1)
	array2 := array1
	fmt.Printf("类型:%T, 地址:%p, 值:%v\n", array2, &array2, array2)
	array1[4] = 20
	fmt.Printf("array1的值:%v\n", array1)
	fmt.Printf("array2的值:%v\n", array2)

	fmt.Println("==========")

	// 切片属于复合类型，存储的是引用地址，传递的是底层引用地址，共享底层数据
	slice3 := array1[:]
	fmt.Printf("类型:%T, 地址:%p, 值:%v\n", slice3, &slice3, slice3)
	slice4 := slice3
	fmt.Printf("类型:%T, 地址:%p, 值:%v\n", slice4, &slice4, slice4)
	for i := 0; i < len(slice3); i++ {
		slice3[i] = (10*3 + 5) + i // 修改slice3切片的值
	}
	fmt.Printf("类型:%T, 地址:%p, 值:%v\n", slice3, &slice3, slice3)
	fmt.Printf("类型:%T, 地址:%p, 值:%v\n", slice4, &slice4, slice4)
}
