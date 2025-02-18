package main

import "fmt"

func main() {
	// map数据结构的基本操作

	// 定义一个map
	m1 := make(map[int]string)

	// 插入/更新键值对
	m1[1] = "Go"
	m1[2] = "Java"
	m1[3] = "C++"
	m1[4] = "Rust"

	// 遍历map
	for k, v := range m1 {
		fmt.Println("Key:", k, "Value:", v)
	}

	// 根据key获取value
	fmt.Println("m1[1]=", m1[1])
	fmt.Println("m1[1]=", m1[2])
	fmt.Println("m1[1]=", m1[3])
	fmt.Println("m1[1]=", m1[4])

	// 删除操作
	delete(m1, 2)

	// 遍历map
	for k, v := range m1 {
		fmt.Println("Key:", k, "Value:", v)
	}
}
