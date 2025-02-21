package main

import "fmt"

// main函数是整个go程序的入口
func main() {
	/*
		函数：function

		定义：一段独立、可重用的代码块，它接收输入（参数），执行特定的任务，并可能返回输出（返回值）。

		语法格式：
			func 函数名(参数列表) (返回值列表) {
				函数体
				// 可选的 return 返回值
			}

			func：定义函数的关键字
			函数名：是函数的标识符，用于调用函数。
			参数列表：定义了函数接收的输入（可以为空）。
			返回值列表：定义了函数返回的结果（可以为空或多个返回值）。
			函数体：是具体的逻辑代码。
			return语句：如果函数声明了返回值，则函数体内必须使用 return 返回对应的值。如果没有返回值类型声明，则可以省略 return。
	*/

	// 功能 10 + 20
	var p1 = 10
	var p2 = 20
	var sum1 = p1 + p2
	fmt.Println(sum1)

	// 功能 30 + 40
	var p3 = 30
	var p4 = 40
	var sum2 = p3 + p4
	fmt.Println(sum2)

	// 思考：要是接下来还是有两个数值进行相加操作，是否还要编写多次
	// 是否可以将其抽取出来处理，实现功能的复用性，这就是函数。

	// 调用函数，传入的参数是实际参数
	res1 := cal(10, 20) // 10 + 20
	fmt.Println(res1)

	res2 := cal(30, 40) // 30 + 40
	fmt.Println(res2)

	res3 := cal(156, 178) // 156 + 178
	fmt.Println(res3)
}

// cal 计算函数，这里函数定义的参数是形式参数
func cal(param1, param2 int) (res int) {
	res = param1 + param2
	return res
}
