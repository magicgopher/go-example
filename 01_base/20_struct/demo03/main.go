package main

import (
	"fmt"
	"github.com/magicgopher/go-example/01_base/20_struct/demo03/mypackage"
)

func main() {
	/*
		结构体可见性
	*/

	s1 := mypackage.Student1{
		Name: "小明",
		Age:  10,
	}
	fmt.Println(s1)

	// 这里student2在其他是不可见的。
	//s2 := mypackage.student2{
	//	name: "小红",
	//	age:  11,
	//}
	//fmt.Println(s2)
}
