package main

import "fmt"

func main() {
	/*
		嵌套结构体
	*/
	p := Person{
		Name: "David",
		Age:  40,
		Addr: Address{
			City:    "Beijing",
			Country: "China",
		},
	}
	fmt.Println(p)           // 输出: {David 40 {Beijing China}}
	fmt.Println(p.Addr.City) // 输出: Beijing
}

type Address struct {
	City    string
	Country string
}

type Person struct {
	Name string
	Age  int
	Addr Address // 嵌套另一个结构体
}
