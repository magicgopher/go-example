package main

import "fmt"

func main() {
	/*
		数据类型：
			A.按照数据类型划分：
				基础类型：int,uint,float,complex,bool,byte,rune,string
				复合类型：array,slice,map,function,struct,pointer,interface,chan

			B.按照数据的存储特点来划分
				值类型的数据：int,uint,float,complex,bool,byte,rune,string,array,struct
				引用类型的数据：slice,map,function,pointer,interface,chan

		函数的参数传递
			值传递：传递的是数据的副本。修改数据，对于原始的数据没有影响的。
				值类型的数据，默认都是值传递：基础类型，array，struct

			引用传递：传递的是数据的地址。导致多个变量指向同一块内存。
				引用类型的数据：默认都是引用传递，slice,map,function,pointer,interface,chan
	*/

	// 定义一个数组
	a1 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("类型：%T, 值:%v\n", a1, a1)
	arrayInfo(a1)
	fmt.Printf("类型：%T, 值:%v\n", a1, a1)

	fmt.Println("===============")

	// 定义一个切片
	s1 := []int{11, 22, 33, 44, 55, 66}
	fmt.Printf("类型：%T, 值:%v\n", s1, s1)
	sliceInfo(s1)
	fmt.Printf("类型：%T, 值:%v\n", s1, s1)
}

// 定义一个函数参数传递数组
func arrayInfo(a [5]int) {
	// 修改传入参数的第一个元素的数据
	a[0] = 999
	fmt.Printf("arrayInfo函数, 类型:%T, 值:%v\n", a, a)
}

// 定义一个函数参数传入切片
func sliceInfo(s []int) {
	// 修改传入参数的第一个元素的数据
	s[0] = 123456
	fmt.Printf("sliceInfo函数, 类型:%T, 值:%v\n", s, s)
}
