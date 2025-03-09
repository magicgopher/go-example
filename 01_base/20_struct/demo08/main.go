package main

import "fmt"

func main() {
	/*
		结构体之间的转换，如果两个结构体的字段完全相同（字段名、字段类型、字段顺序都一致），它们可以直接进行类型转换。
		这种转换本质上是类型断言，Go 会将一个结构体的内存布局直接解释为另一个结构体。
	*/

	p := Person{"小明", 17}
	e := Employee(p) // 直接转换
	fmt.Println("Person:", p)
	fmt.Println("Employee:", e)
	//t := Teacher(p) // 编译错误

	//s := Student(p) // 错误：不能直接转换
	s := Student{
		Name:  p.Name, // 手动赋值相同字段
		Grade: 10,     // 为不同字段提供值
	}
	println(s.Name, s.Grade) // 输出: Bob 10
}

type Person struct {
	Name string
	Age  uint
}

type Employee struct {
	Name string
	Age  uint
}

type Student struct {
	Name  string
	Grade uint
}

type Teacher struct {
	name string
	age  uint
}
