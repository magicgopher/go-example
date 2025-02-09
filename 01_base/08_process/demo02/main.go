package main

import "fmt"

func main() {
	// if语句：表示判断结果为true执行的语句体。
	/*
		语法格式：
			if 条件表达式 {
				语句体
			}
	*/

	// 给定一个数字判断是否大于20,打印大于20,否则不打印
	num := 16
	if num > 20 {
		fmt.Println("大于20")
	}
	// 给定一个成绩判断是否及格,及格则打印成绩及格
	score := 80
	if score >= 60 {
		fmt.Println("成绩及格.")
	}
	fmt.Println("main over")
}
