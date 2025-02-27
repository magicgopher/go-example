package main

import "fmt"

func main() {
	/*
		高阶函数：是指可以将函数作为参数传递给另一个函数，或者返回一个函数作为结果的函数。

		fun1(), fun2()
		将fun1函数作为了fun2这个函数的参数。

			fun2函数：就叫高阶函数
				接收了一个函数作为参数的函数，高阶函数

			fun1函数：回调函数
				作为另一个函数的参数的函数，叫做回调函数。
	*/

	// 设计一个函数，用于求两个整数的加减乘除运算
	fmt.Printf("%T\n", add)      // func(int, int) int
	fmt.Printf("%T\n", operator) // func(int, int, func(int, int) int) int

	// add
	res1 := add(10, 20)
	fmt.Println(res1) // 30

	// sub
	res2 := sub(20, 5)
	fmt.Println(res2) // 25

	// add
	res3 := operator(200, 20, add)
	fmt.Println(res3)

	// sub
	res4 := operator(200, 20, sub)
	fmt.Println(res4)

	// mul
	fun1 := func(a, b int) int {
		return a * b
	}
	res5 := operator(6, 6, fun1)
	fmt.Println(res5) // 36

	// div
	fun2 := func(a, b int) int {
		return a / b
	}
	res6 := operator(77, 7, fun2)
	fmt.Println(res6) // 11

	// 调用apply高阶函数
	res7 := apply([]int{1, 2, 3, 4}, func(i int) int {
		return i * 3
	})
	fmt.Println(res7) // [3,6,9,12]

	// 调用multiplier函数
	double := multiplier(2) // 返回一个将输入乘以2的函数
	triple := multiplier(3) // 返回一个将输入乘以3的函数

	fmt.Println(double(5)) // 输出: 10
	fmt.Println(triple(5)) // 输出: 15
}

// operator 运算
func operator(a, b int, fun func(int, int) int) int {
	fmt.Println(a, b, fun) // 打印三个参数
	result := fun(a, b)
	return result
}

// add
func add(a, b int) int {
	return a + b
}

// sub
func sub(a, b int) int {
	return a - b
}

// 定义一个高阶函数，接受一个函数作为参数
func apply(slice []int, fn func(int) int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// 返回一个函数
func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
