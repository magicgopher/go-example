package main

import "fmt"

func main() {
	// 顺序结构
	// 执行顺序确实是从上到下，但具体到每一行代码，执行顺序是从左到右的。
	fmt.Println("这是第一条语句")
	fmt.Println("这是第二条语句")
	fmt.Println("这是第三条语句")

	a := 1
	b := 2
	c := a + b // 先计算 a + b，然后赋值给 c
	fmt.Println(c)
}
