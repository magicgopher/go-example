package main

import "fmt"

func main() {
	/*
		从现有的数组转为切片
		语法格式：slice := array[start:end]
		例如：s1 := arr[0:9] 这里从索引0开始（包括）到索引9（不包括）
	*/

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("----------1.已有数组直接创建切片--------------------")
	s1 := a[:5]  //1-5
	s2 := a[3:8] //4-8
	s3 := a[5:]  // 6-10
	s4 := a[:]   // 1-10
	fmt.Println("a:", a)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)

	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", s1)

	fmt.Println("----------2.长度和容量--------------------")
	fmt.Printf("s1	len:%d,cap:%d\n", len(s1), cap(s1)) //s1	len:5,cap:10
	fmt.Printf("s2	len:%d,cap:%d\n", len(s2), cap(s2)) //s2	len:5,cap:7
	fmt.Printf("s3	len:%d,cap:%d\n", len(s3), cap(s3)) //s3	len:5,cap:5
	fmt.Printf("s4	len:%d,cap:%d\n", len(s4), cap(s4)) //s4	len:10,cap:10

	fmt.Println("----------3.更改数组的内容--------------------")
	a[4] = 100
	fmt.Println(a)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	fmt.Println("----------4.更改切片的内容--------------------")
	s2[2] = 200
	fmt.Println(a)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	fmt.Println("----------4.更改切片的内容--------------------")
	s1 = append(s1, 1, 1, 1, 1)
	fmt.Println(a)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println("----------5.添加元素切片扩容--------------------")
	fmt.Println(len(s1), cap(s1))

	s1 = append(s1, 2, 2, 2, 2, 2)
	fmt.Println(a)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(len(s1), cap(s1))
	fmt.Printf("%p\n", s1)
	fmt.Printf("%p\n", &a)
}
