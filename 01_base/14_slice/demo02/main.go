package main

import "fmt"

func main() {
	/*
		切片的扩容机制
			当切片容量小于1024时：新容量 = 旧容量 * 2
			当切片容量大于等于1024时：新容量 = 旧容量 * 1.25 (每次增加25%)
	*/

	// 初始切片
	s := make([]int, 0)
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))

	// 添加元素，观察容量变化
	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
	}
}
