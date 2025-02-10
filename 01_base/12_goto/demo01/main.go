package main

import "fmt"

func main() {
	// goto关键字：用于将程序控制权转移到指定的标签处。简单来说，goto 语句可以跳转到程序中标记的某个位置继续执行代码。
	// 语法格式： goto label
	// goto 关键字的注意事项
	// 1.goto 语句只能在同一个函数内部跳转，不能跨函数跳转。
	// 2.goto 语句不能跳转到循环语句内部，但是可以从循环语句内部跳转到循环语句外部。
	// 3.goto 语句不能跳转到闭包内部。
	// 4.goto 语句的使用应该谨慎，过度的使用可能会导致代码可读性降低，增加代码维护难度。
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i*j > 30 {
				goto breakLoop // 跳出多层循环
			}
			fmt.Println(i, j)
		}
	}
breakLoop:
	fmt.Println("跳出循环")
}
