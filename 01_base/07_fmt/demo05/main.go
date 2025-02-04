package main

import "fmt"

func main() {
	// fmt 包（format）的缩写
	// fmt.Scanf 函数类似于 C 语言中的 scanf 函数，它可以根据指定的格式化字符串来读取输入。
	var name string
	var age int

	fmt.Print("请输入姓名和年龄（用逗号分隔）：")
	fmt.Scanf("%s,%d", &name, &age) // 注意：这里使用了逗号分隔

	fmt.Printf("姓名：%s，年龄：%d\n", name, age)
}
