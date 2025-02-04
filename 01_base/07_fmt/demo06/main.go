package main

import "fmt"

func main() {
	// fmt 包（format）的缩写
	// fmt.Scanln 函数用于读取一行输入，并以换行符作为分隔符。
	var name string
	var age int

	fmt.Print("请输入姓名和年龄（用空格或换行分隔）：")
	fmt.Scanln(&name, &age) // 注意：这里使用了 Scanln

	fmt.Printf("姓名：%s，年龄：%d\n", name, age)
}
