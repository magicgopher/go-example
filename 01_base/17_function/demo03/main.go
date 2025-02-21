package main

import "fmt"

func main() {
	/*
		函数的参数：
			形式参数：也叫形参。函数定义的时候，用于接收外部传入的数据的变量。
				函数中，某些变量的数值无法确定，需要由外部传入数据。

			实际参数：也叫实参。函数调用的时候，给形参赋值的实际的数据


		函数调用：
			1.函数名：声明的函数名和调用的函数名要统一
			2.实参必须严格匹配形参：顺序，个数，类型，一一对应的。

	*/

	// 调用函数 这里传入的 1, 100 就是实际参数
	res := sum(1, 100)
	fmt.Println(res)
}

// 定义一个求和函数 start, end 参数是形式参数
func sum(start, end int) (res int) {
	for i := start; i <= end; i++ {
		res += i
	}
	return res
}
