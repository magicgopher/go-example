package main

import "fmt"

func main() {
	/*
		golang中循环只有for循环
			for循环语法格式1：
				for 初始化条件; 判断条件; 后置语句 {
					// 循环体的代码部分
				}

			for循环语法格式2：
				初始化条件
				for 判断条件 {
					// 循环体的代码部分
					后置语句
				}

			for死循环语法格式:
				for {
					// 循环体的代码部分
				}
	*/

	// 循环打印10次hello world
	//for i := 1; i <= 10; i++ {
	//	fmt.Println("hello world!")
	//}
	fmt.Println("------------------")

	i := 1
	for i <= 10 {
		fmt.Println("hello world")
		i++
	}
	fmt.Println("------------------")
}
