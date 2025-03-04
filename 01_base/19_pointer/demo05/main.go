package main

import "fmt"

func main() {
	/*
		函数指针：一个指针，指向一个函数的指针
			function作为复合类型（派生）类型，默认看作一个指针，没有 *

		指针函数：一个函数，该函数的返回值是一个指针
	*/

	// 定义函数变量
	var f1 func()
	fmt.Println(f1) // nil
	f1 = fun1
	fmt.Println(f1)
	f1() // 调用这个函数变量，f1指向fun1函数地址，所以使用()就是调用fun1函数

	fmt.Println("===============")

	arr1 := fun2()
	fmt.Printf("arr1的类型:%T, 地址:%p, 值:%v\n", arr1, &arr1, arr1)

	fmt.Println("===============")

	arr2 := fun3()
	fmt.Printf("arr2的类型:%T，地址:%p，值:%v\n", arr2, &arr2, arr2)
	fmt.Printf("arr2指针中存储的数组的地址:%p\n", arr2)

	fmt.Println("===============")

	// 指针函数
	res1 := newInt(10)
	fmt.Println(res1)  // 输出的是指针变量指向的地址值
	fmt.Println(*res1) // 输出的指针指向的地址存储的值
}

// newInt
func newInt(value int) *int {
	v := value
	return &v
}

// fun3
func fun3() *[3]int {
	arr := [3]int{10, 20, 30}
	fmt.Printf("fun3函数中arr的地址:%p\n", &arr)
	return &arr
}

// fun2 普通函数 返回值是一个数组
func fun2() [3]int {
	return [3]int{1, 2, 3}
}

func fun1() {
	fmt.Println("fun1()")
}
