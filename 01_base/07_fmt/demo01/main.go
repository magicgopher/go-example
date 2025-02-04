package main

import "fmt"

func main() {
	// fmt 包是一个非常重要的标准库，它提供了一系列函数，用于格式化输入和输出。

	name := "MagicGopher"
	age := 20

	// Print函数：用于将数据格式化并输出到标准输出（通常是控制台）。
	n, err := fmt.Print("姓名:", name, "年龄:", age)
	if err != nil {
		fmt.Println("输出时发生错误：", err)
	} else {
		fmt.Println("成功写入字节数：", n)
	}
}
