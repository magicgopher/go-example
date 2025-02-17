package main

import "fmt"

func main() {
	// 使用 for range 遍历切片
	slice1 := []int{10, 20, 30, 40, 50}
	for i, v := range slice1 {
		fmt.Printf("索引:%d, 值:%d\n", i, v)
	}
}
