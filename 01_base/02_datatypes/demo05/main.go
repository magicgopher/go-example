package main

import "fmt"

func main() {
	// 字符
	var a byte = 'A' // 等同于 uint8
	var b rune = '本' // 等同于 int32
	fmt.Printf("%c %c\n", a, b)

	// 字符串
	var s string = "Hello, World!"
	fmt.Println(s)
}
