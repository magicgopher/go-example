package main

import "fmt"

func main() {
	// 变量：计算机的一小块内存，用于存储数据，在程序运行的过程中，数据是可变的。

	// 变量定义的第一种方式：声明变量赋值
	// 语法格式：var 变量名称 = 赋值的数据
	var a int // 分开写
	a = 100
	var b int = 200 // 写在一行
	fmt.Printf("变量a的数据类型是:%T,值是:%d\n", a, a)
	fmt.Printf("变量b的数据类型是:%T,值是:%d\n", b, b)

	// 变量定义的第二种方式：简短定义
	// 语法格式：变量名称 := 赋值的数据
	name := "golang gopher"
	fmt.Printf("name变量的数据类型是%T,值是:%v\n", name, name)

	// 变量定义的第三种方式：类型推断
	// 语法格式：var 变量名称 = 赋值的数据
	var flag = 8.88 // 根据赋值的数据类型来推断变量的数据类型
	fmt.Printf("flag变量的数据类型:%T,值是:%v\n", flag, flag)
}
