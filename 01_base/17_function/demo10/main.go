package main

import "fmt"

func main() {
	/*
		defer：这个英文有延迟、延缓的意思。
		在go中通过defer关键字来延迟执行一个函数（一个方法）或者语句。

		语法格式：
			defer 延迟的函数或者语句

		defer关键字遵循“后进先出”（LIFO）的原则，这意味着后defer的先执行，先defer的后执行。

	*/

	// 执行打印结果是3 2 1
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}
