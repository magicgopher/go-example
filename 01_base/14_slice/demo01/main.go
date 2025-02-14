package main

import "fmt"

func main() {
	// 数组：固定长度，存储一种类型的数据结构
	// 切片slice：类似数组的数据结构，但是可以动态地调整大小，是一个引用类型的容器，指向了一个底层数组。
	// 切片slice的定义：[]类型
	// 使用make函数创建切片 func make(t Type, size ...IntegerType) Type
	// 创建切片时，make 函数接受三个参数：
	// 第一个参数：表示类型这里有slice切片、map键值对、chan通道
	// 第二个参数：表示长度
	// 第三个参数：表示容量

	// 创建一个切片
	var s1 []int
	// 打印输出切片的值
	fmt.Printf("s1切片的值:%v\n", s1)
	// 打印切片的长度和容量
	fmt.Printf("s1切片的长度:%d, 容量:%d\n", len(s1), cap(s1))

	// 使用 append 函数给切片的末尾添加元素
	// 必须接收返回值： 由于 append 函数可能会返回一个新的切片，因此必须使用原切片变量接收返回值，否则可能会导致数据丢失。
	s1 = append(s1, 10, 20, 30, 40, 50)
	// 打印输出切片的值
	fmt.Printf("s1切片的值:%v\n", s1)
	// 打印切片的长度和容量
	fmt.Printf("s1切片的长度:%d, 容量:%d\n", len(s1), cap(s1))

}
