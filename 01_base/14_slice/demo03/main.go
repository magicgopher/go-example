package main

import (
	"fmt"
)

func main() {
	/*
		切片Slice：
				1.每一个切片引用了一个底层数组
				2.切片本身不存储任何数据，都是这个底层数组存储，所以修改切片也就是修改这个数组中的数据
				3.当向切片中添加数据时，如果没有超过容量，直接添加，如果超过容量，自动扩容(成倍增长)
				4.切片一旦扩容，就是重新指向一个新的底层数组

		s1的cap容量从 3 -> 6 -> 12 -> 24
	*/
	s1 := []int{1, 2, 3}
	sliceInfo(s1)

	s1 = append(s1, 4, 5)
	sliceInfo(s1)

	s1 = append(s1, 6, 7, 8)
	sliceInfo(s1)

	s1 = append(s1, 9, 10)
	sliceInfo(s1)

	s1 = append(s1, 11, 12, 13, 14, 15)
	sliceInfo(s1)
}

func sliceInfo(s []int) {
	// 打印数组信息
	fmt.Printf("len:%d,cap:%d\n", len(s), cap(s))
	fmt.Printf("数组: %v\n", s)
	fmt.Printf("%p\n", s)
	fmt.Println("==========")
}
