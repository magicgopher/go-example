package main

import "fmt"

func main() {
	/*
		go基本类型、array数组、struct结构体属于值类型

		function也属于引用类型
	*/

	a := 10
	fmt.Printf("类型:%T\n", a)

	b := [3]int{1, 2, 3}
	fmt.Printf("类型:%T\n", b)

	c := make(map[int]string)
	fmt.Printf("类型:%T\n", c)

	fmt.Printf("类型:%T\n", fun1) // 类型:func()
	fmt.Printf("类型:%T\n", fun2) // 类型:func(int) int
	fmt.Printf("类型:%T\n", fun3) // 类型:func(float64, int, int) (int, int)
	fmt.Printf("类型:%T\n", fun4) // 类型:func(string, string, int, int) (string, int, float64)
}

// fun1
func fun1() {

}

// fun2
func fun2(a int) int {
	return 0
}

// fun3
func fun3(a float64, b, c int) (int, int) {
	return 0, 0
}

// fun4
func fun4(a, b string, c, d int) (string, int, float64) {
	return "", 0, 0
}
