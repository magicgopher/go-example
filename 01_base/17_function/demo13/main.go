package main

import "fmt"

func main() {
	/*
			匿名函数：是指没有函数名的函数定义。它通常被用来定义一个临时的、一次性使用的函数。

			匿名函数可以直接赋值给变量，或者作为参数传递给其他函数，甚至可以立即调用。

			匿名函数定义：
				func(参数列表) 返回值类型 {
		    		// 函数体
				}(传入的参数列表)

			匿名函数：
				Go语言是支持函数式编程：
				1.将匿名函数作为另一个函数的参数，回调函数
				2.将匿名函数作为另一个函数的返回值，可以形成闭包结构。
	*/

	// 定义一个匿名函数，入参数是一个string类型返回值也是一个string类型，将返回值赋值给res1变量
	result1 := func(s string) string {
		return fmt.Sprintf("%s world!", s)
	}("hello")
	fmt.Println(result1)

	// 立即执行的匿名函数
	result2 := func(x int) int {
		return x + 3
	}(7)
	fmt.Println(result2) // 输出 10
}
