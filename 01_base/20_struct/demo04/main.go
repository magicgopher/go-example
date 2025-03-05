package main

import "fmt"

func main() {
	/*
		值类型是传递数据的副本，在 Go 中值类型有基础类型(int, uint, float, bool, string等)、array、struct

		引用类型是传递内存指向的地址，在 Go 中引用类型有slice、map、function、pointer、chan、interface

		结构体struct也是值类型
	*/

	// 创建一个结构体
	p1 := Person{"小明", 17, "男"}
	fmt.Println(p1)
	fmt.Printf("地址:%p, 类型:%T\n", &p1, p1)

	p2 := p1
	fmt.Println(p2)
	fmt.Printf("地址:%p, 类型:%T\n", &p2, p2)

	p2.name = "小红"
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println("===============")

	// 定义结构体指针
	var ptr1 *Person
	ptr1 = &p1
	fmt.Println(ptr1)
	fmt.Printf("自身地址:%p, 指向的地址:%p, 类型:%T, 值:%v\n", &ptr1, ptr1, ptr1, ptr1)

	ptr1.name = "张三"  // 这操作的是指针变量指向的内存地址，也就是p1的地址存储的数据
	fmt.Println(ptr1) // &{张三 17 男}
	fmt.Println(p1)   // {张三 17 男}

	fmt.Println("===============")

	// func new(Type) *Type 作用：指定类型分配内存，并返回指向该内存的指针。
	ptr2 := new(Person)
	fmt.Printf("类型:%T\n", ptr2)
	fmt.Println(ptr2)
	ptr2.name = "李四"
	ptr2.age = 20
	ptr2.sex = "女"
	fmt.Println(ptr2)

	fmt.Println("===============")

	ptr3 := new(string) // 返回的是该类型的内存指针
	fmt.Printf("类型:%T, 值:%v\n", ptr3, ptr3)
}

type Person struct {
	name string
	age  uint
	sex  string
}
