package main

import "fmt"

func main() {
	/*
		map遍历对比slice遍历

		使用 for range 进行遍历 map、slice

		slice：index, value
		map：key, value
	*/

	// map
	m1 := map[int]string{
		1: "张三",
		2: "李四",
		3: "王五",
		4: "老六",
		5: "小七",
		6: "老八",
	}
	// 遍历map
	// go map 数据结构是无序的。
	// 这意味着你不能依赖于 map 中键值对的任何特定顺序。每次遍历 map，键值对的顺序都可能不同。
	for k, v := range m1 {
		fmt.Println("k:", k, "v:", v)
	}

	fmt.Println("==========")

	// slice
	s1 := []int{98, 66, 31, 28, 86}
	// 遍历slice
	for i, v := range s1 {
		fmt.Println("i:", i, "v:", v)
	}
}
