package main

import "fmt"

func main() {
	y := 10
	for x := 5; x < 20; x++ {
		switch x {
		case y + 10:
			fmt.Printf("x (%d) 等于 y + 10 (%d)\n", x, y+10)
		case y * 2:
			fmt.Printf("x (%d) 等于 y * 2 (%d)\n", x, y*2)
		default:
			fmt.Printf("x (%d) 不等于 y + 10 或 y * 2\n", x)
		}
	}

	/*
		代码解释：

		定义变量 y： 首先定义一个整数变量 y，并赋值为 10。
		循环遍历 x： 使用 for 循环，让 x 从 5 遍历到 19。
		switch 语句： 在每次循环中，使用 switch 语句检查 x 的值。
		case 语句：
			第一个 case 语句检查 x 是否等于 y + 10。
			第二个 case 语句检查 x 是否等于 y * 2。
		default 语句： 如果 x 不等于任何一个 case 语句中的值，则执行 default 语句中的代码。
		打印输出： 根据 x 的值，打印相应的输出信息。
	*/
}
