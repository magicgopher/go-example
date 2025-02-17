package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{6, 7, 8}
	fmt.Println("拷贝前...")
	fmt.Printf("长度:%d, 容量:%d, 值:%v\n", len(s1), cap(s1), s1) // 1 2 3 4 5
	fmt.Printf("长度:%d, 容量:%d, 值:%v\n", len(s2), cap(s2), s2) // 6 7 8

	copy(s1, s2[:]) // 将 s2 所有元素拷贝到 s1 中
	fmt.Println("拷贝后...")
	fmt.Printf("长度:%d, 容量:%d, 值:%v\n", len(s1), cap(s1), s1) // 6 7 8 4 5
	fmt.Printf("长度:%d, 容量:%d, 值:%v\n", len(s2), cap(s2), s2) // 6 7 8

	fmt.Println("==========")

	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s4 := []int{11, 12, 13, 14}
	fmt.Println("拷贝前...")
	fmt.Printf("长度:%d, 容量:%d, 值:%v\n", len(s3), cap(s3), s3) // 1-10
	fmt.Printf("长度:%d, 容量:%d, 值:%v\n", len(s4), cap(s4), s4) // 11-14

	copy(s3[3:7], s4[:])
	fmt.Println("拷贝后...")
	fmt.Printf("长度:%d, 容量:%d, 值:%v\n", len(s3), cap(s3), s3) // 1 2 3 11 12 13 14 8 9 10
	fmt.Printf("长度:%d, 容量:%d, 值:%v\n", len(s4), cap(s4), s4) // 11-14

	fmt.Println("==========")

	s5 := []int{21, 22, 23, 24, 25}
	s6 := []int{31, 32, 33, 34, 35}
	fmt.Println("拷贝前...")
	fmt.Printf("长度:%d, 容量:%d, 值:%v\n", len(s5), cap(s5), s5) // 21 22 23 24 25
	fmt.Printf("长度:%d, 容量:%d, 值:%v\n", len(s6), cap(s6), s6) // 31 32 33 34 35

	copy(s5, s6[:2]) // 拷贝 s6 切片索引0到索引2的元素（不包含索引2的元素）到切片s5
	fmt.Println("拷贝后...")
	fmt.Printf("长度:%d, 容量:%d, 值:%v\n", len(s5), cap(s5), s5) // 31 32 23 24 25
	fmt.Printf("长度:%d, 容量:%d, 值:%v\n", len(s6), cap(s6), s6) // 31 32 33 34 35
}
