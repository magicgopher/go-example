package main

import "fmt"

func main() {
	// 同时定义多个变量
	var a, b, c = 100, 8.88, true
	fmt.Println(a, b, c)

	// 使用简短定义多个变量，必须要有一个新的变量
	//a,b,c := 200, 9.99, false
	a, b, c, d := 200, 9.99, false, "hello"
	fmt.Println(a, b, c, d)

	// 定义多个同一个数据类型的变量
	var f, g, h int
	f = 11
	g = 22
	h = 33
	fmt.Println(f, g, h)

	// 使用()方式定义多个变量
	var (
		name    = "gopher"
		age     = 20
		address = "地球"
	)
	fmt.Printf("name: %s, age: %d, address: %s\n", name, age, address)
}
