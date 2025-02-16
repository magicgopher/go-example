package main

import "fmt"

func main() {
	/*
		深拷贝：拷贝的是数据本身。
				值类型的数据，默认都是深拷贝：array，int，float，string，bool，struct


			浅拷贝：拷贝的是数据 地址。
				导致多个变量指向同一块内存
				引用类型的数据，默认都是浅拷贝：slice，map，

				因为切片是引用类型的数据，直接拷贝的是地址。

			func copy(dst, src []Type) int
				dst：目标切片，数据将被复制到这个切片。
				src：源切片，数据将从这个切片复制。
				Type：切片中元素的类型，dst 和 src 的元素类型必须相同。
				返回值: 复制的元素数量，即实际复制的字节数。
				可以实现切片的拷贝
	*/

	s1 := []int{1, 2, 3, 4}
	s2 := make([]int, 0) // len:0, cap:0
	// 使用循环的方式将s1切片的元素赋值给s2切片
	for i := 0; i < len(s1); i++ {
		s2 = append(s2, s1[i])
	}
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

	fmt.Println("==========")

	s1[0] = 100
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

	fmt.Println("==========")

	// 使用copy()函数将s1切片的元素拷贝给s2切片
	s3 := []int{7, 8, 9, 10}
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)

	fmt.Println("==========")

	copy(s3, s2[:2]) // 将 s2 切片的[:2]从0索引到索引2（不包括2）的元素，拷贝到 s3 切片
	fmt.Printf("长度:%d, 容量:%d, 值:%v\n", len(s2), cap(s2), s2)
	fmt.Printf("长度:%d, 容量:%d, 值:%v\n", len(s3), cap(s3), s3)

	fmt.Println("==========")

	s4 := []int{11, 12, 13, 14, 15, 16}
	s5 := []int{17, 18, 19, 20}
	fmt.Printf("长度:%d, 容量:%d, 值:%v\n", len(s4), cap(s4), s4)
	fmt.Printf("长度:%d, 容量:%d, 值:%v\n", len(s5), cap(s5), s5)
	copy(s4, s5[:]) // 将 s5 切片的全部元素拷贝到 s4 切片
	fmt.Printf("长度:%d, 容量:%d, 值:%v\n", len(s4), cap(s4), s4)
	fmt.Printf("长度:%d, 容量:%d, 值:%v\n", len(s5), cap(s5), s5)
}
