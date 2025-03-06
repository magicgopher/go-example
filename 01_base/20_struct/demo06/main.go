package main

import "fmt"

func main() {
	/*
		嵌套结构体：嵌套结构体是指一个结构体包含另一个结构体作为字段。
		嵌套可以是命名字段（显式命名）或匿名字段（无显式名称，直接嵌入类型）。
		匿名字段是嵌套结构体的一种特殊形式，但它的行为和使用方式与普通命名字段的嵌套有所不同。
	*/

	p := Person{
		Name: "Alice",
		Age:  25,
		Sex:  "男性",
		Addr: Address{ // Addr是值传递
			City:    "Beijing",
			Country: "China",
		},
	}

	// 访问字段必须通过字段名
	fmt.Println("Person:", p)
	fmt.Println("City:", p.Addr.City)
	fmt.Println("Country:", p.Addr.Country)

	p1 := Person1{
		"Alan",
		20,
		"女性",
		&Address{"Tokyo", "Japan"}, // 引用传递
	}
	fmt.Println(p1)
	fmt.Println(p1)
}

// Person 人
type Person struct {
	Name string  // 姓名
	Age  uint    // 年龄
	Sex  string  // 性别
	Addr Address // 地址
}

// Address 地址
type Address struct {
	City    string // 城市
	Country string // 国家
}

// Person1 p1
type Person1 struct {
	Name string   // 姓名
	Age  uint     // 年龄
	Sex  string   // 性别
	Addr *Address // 地址（指针类型）
}
