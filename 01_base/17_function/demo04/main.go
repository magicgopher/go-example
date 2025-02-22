package main

import "fmt"

func main() {
	/*
		可变参数：
					概念：一个函数的参数的类型确定，但是个数不确定，就可以使用可变参数。

					语法：
						func functionName(param ...Type) {
							// 函数体
						}

						对于函数，可变参数相当于一个切片。
						调用函数的时候，可以传入0-多个参数。

						Println(),Printf(),Print()
						append()

					注意事项：
						A：如果一个函数的参数是可变参数，同时还有其他的参数，可变参数要放在
							参数列表的最后。
						B：一个函数的参数列表中最多只能有一个可变参数。
	*/

	// 调用函数
	res1 := sum(1, 2, 3)
	fmt.Println(res1)

	res2 := sum(1, 2, 3, 4, 5)
	fmt.Println(res2)
}

// sum 求和（这里参数定义为可变参数）
func sum(params ...int) (res int) {
	for _, v := range params {
		res += v
	}
	return res
}
