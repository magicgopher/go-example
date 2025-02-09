package main

import "fmt"

func main() {
	/*
		switch中的break和fallthrough语句

		switch expression {
		case value1:
			// value1 匹配时执行的代码
		case value2:
			// value2 匹配时执行的代码
		default:
			// 没有匹配的 case 时执行的代码
		}

		break：可以使用在switch中，也可以使用在for循环中
			强制结束case语句，从而结束switch分支

		fallthrough：用于穿透switch
			当switch中某个case匹配成功之后，就执行该case语句
			如果遇到fallthrough，那么后面紧邻的case，无需匹配， 执行穿透执行。

			fallthrough应该位于某个case的最后一行
	*/

	// 判断星期几
	num := 5
	switch num {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	case 6:
		fmt.Println("星期六")
	case 7:
		fmt.Println("星期天")
	default:
		fmt.Println("日期数值有误。")
	}
}
