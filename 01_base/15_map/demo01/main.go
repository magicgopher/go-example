package main

import "fmt"

func main() {
	/*
		map：是一种存储键值对（key-value）的数据结构。

		语法格式：map[键的数据类型]值的数据类型
	*/

	// 定义map的三种方式

	// 第一种
	var m1 map[int]string     // 此时创建的map是一个nil，这意味着它没有指向任何底层的数据结构，因此不能直接进行赋值操作。
	fmt.Println(m1 == nil)    // true
	m1 = make(map[int]string) // 初始化map
	fmt.Println(m1 == nil)    // false
	fmt.Println(m1)           // map[]

	// 第二种
	m2 := make(map[string]string)
	fmt.Println(m2 == nil)
	fmt.Println(m2) // map[]

	// 第三种
	m3 := map[string]string{"Go": "Fiber", "Java": "Spring", "C++": "QT"}
	fmt.Println(m3 == nil) // false
	fmt.Println(m3)        // map[C++:QT Go:Fiber Java:Spring]
}
