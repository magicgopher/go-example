package main

import "fmt"

func main() {
	// fmt 包中的格式化字符串函数

	// Sprint
	str1 := fmt.Sprint("Hello", "World") // 返回: "Hello World"
	fmt.Printf("类型:%T, 值:%s\n", str1, str1)

	// Sprintln
	str2 := fmt.Sprintln("Hello", "World") // 返回："Hello World"
	fmt.Printf("类型:%T, 值:%s\n", str2, str2)

	// Sprintf
	str3 := fmt.Sprintf("Name: %s, Age: %d", "Alice", 25) // 返回: "Name: Alice, Age: 25"
	fmt.Print(str3)
}
