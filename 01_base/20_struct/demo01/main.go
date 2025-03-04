package main

import "fmt"

func main() {
	/*
		结构体：是一种用户自定义的复合数据类型，用于将多个不同类型的数据字段（field）组合成一个单一的逻辑单元。
	*/

	// 创建结构体
	// 方式一
	var u1 user
	fmt.Println("user1:", u1)
	u1.name = "张三"
	u1.age = 18
	u1.sex = "男"
	u1.address = "上海"
	fmt.Println("user1:", u1)

	// 方式二
	u2 := user{
		name:    "李四",
		age:     20,
		sex:     "女",
		address: "深圳",
	}
	fmt.Println("user2:", u2)

	// 方式三
	u3 := user{}
	u3.name = "王五"
	u3.age = 22
	u3.sex = "男"
	u3.address = "广州"
	fmt.Println("user3:", u3)

	// 方式四
	u4 := user{"赵六", 17, "女", "北京"}
	fmt.Println("user4:", u4)
}

// User 定义结构体
type user struct {
	name    string
	age     int
	sex     string
	address string
}
