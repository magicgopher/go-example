package main

import (
	"fmt"
)

func main() {
	// fmt 包是一个非常重要的标准库，它提供了一系列函数，用于格式化输入和输出。

	// Printf函数：格式化输出函数。它允许你通过使用格式化字符串来精确控制输出的格式和内容。

	name := "MagicGopher"
	age := 20
	address := "地球"
	// %v 是默认格式表示。
	fmt.Printf("name:%v, age:%v, address:%v\n", name, age, address)

	// %+v: 类似%v，但输出结构体时会添加字段名。
	u1 := User{name: "张三", age: 19}
	u2 := User{}
	fmt.Printf("user1:%+v\n", u1)
	fmt.Printf("user2:%+v\n", u2)

	// %#v: 值的Go语法表示。
	fmt.Printf("user1:%#v\n", u1)
	fmt.Printf("user2:%#v\n", u2)

	// %t：是布尔类型格式化占位符
	fmt.Printf("布尔类型:%t\n", true)

	// %s：是字符串格式化占位符
	fmt.Printf("字符串类型:%s\n", "hello golang")

	var a = 101
	// %b：是二进制格式化占位符
	fmt.Printf("二进制:%b\n", a)

	// %d：是十进制格式化占位符
	fmt.Printf("十进制:%d\n", a)

	// %o：是八进制格式化占位符
	fmt.Printf("八进制:%o\n", a)

	// %x：是十六进制格式化占位符（小写字母）
	fmt.Printf("十六进制:%x\n", a)

	// %X：是十六进制格式化占位符（大写字母）
	fmt.Printf("十六进制:%X\n", a)

	// %f: 浮点数占位符。.2 表示保留小数点后两位。
	fmt.Printf("浮点数:%.2f\n", 8.8888)

	// %T：它表示输出值的类型。
	fmt.Printf("数据类型:%T\n", 100)

	// %p：指针地址表示。
	var f float64 = 56.99
	fmt.Printf("地址:%p\n", &f)
}

// User 用户
type User struct {
	name string
	age  int
}
