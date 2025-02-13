package main

import "fmt"

func main() {
	// 使用 for range 遍历数组
	arr1 := [5]int{19, 22, 36, 42, 57}
	// for 循环遍历
	for i := 0; i <= len(arr1)-1; i++ {
		fmt.Printf("arr1[%d]=%d\n", i, arr1[i])
	}
	fmt.Println("==========")
	// for range遍历
	for index, value := range arr1 {
		fmt.Printf("arr1[%d]=%d\n", index, value)
	}
}
