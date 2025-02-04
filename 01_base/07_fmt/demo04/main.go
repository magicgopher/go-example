package main

import "fmt"

func main() {
	// fmt 包（format）的缩写
	// fmt.Scan 函数用于从标准输入（通常是键盘）读取数据。它的返回值 n 表示成功读取并赋值给变量的参数个数。
	var name string
	var age int

	fmt.Print("请输入姓名和年龄（用空格分隔）：")
	n, err := fmt.Scan(&name, &age)

	if err != nil {
		fmt.Println("输入有误：", err)
	} else {
		fmt.Printf("成功读取了 %d 个参数\n", n)
		fmt.Printf("姓名：%s，年龄：%d\n", name, age)
	}
}
