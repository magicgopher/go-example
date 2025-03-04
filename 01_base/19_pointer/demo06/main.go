package main

import (
	"fmt"
)

func main() {
	/*
		指针作为函数的参数传递。

		函数参数的传递：值传递，引用传递
	*/

	a := 10
	fmt.Println("调用fun1()函数前, a的值:", a) // a = 10
	fun1(a)
	fmt.Println("调用fun1()函数后，a的值:", a) // a = 10

	fmt.Println("===============")

	b := 20
	fmt.Println("调用fun2()函数前，b的值:", b) // b = 20
	fun2(&b)
	fmt.Println("调用fun2()函数后，b的值:", b) // b = 222

	fmt.Println("===============")

	c := [3]int{1, 2, 3}
	fmt.Println("调用fun3()函数前，c的值:", c) // c = [1,2,3]
	fun3(c)
	fmt.Println("调用fun3()函数后，c的值:", c) // c = [100,2,3]

	fmt.Println("===============")

	d := [3]int{1, 2, 3}
	fmt.Println("调用fun4()函数前，d的值:", d) // d = [1,2,3]
	fun4(&d)
	fmt.Println("调用fun4()函数后，d的值:", d) // d = [4,5,6]
}

// fun1 传递值类型参数
func fun1(a int) { // 值类型传递
	fmt.Println("传入的参数a的值:", a)
	a = 100 // 修改了a的值，但是fun1函数中的a是一个值，不是指针，所以fun1函数中的a不会改变
	fmt.Println("fun1()函数中修改了a的值，a的值:", a)
}

// fun2 传递指针类型参数
func fun2(b *int) { // 引用类型传递
	fmt.Println("传入的参数b的值:", *b)
	*b = 222 // 修改指针指向的地址存储的值
	fmt.Println("fun2()函数中修改了b的值，b的值:", *b)
}

// fun3 传递数组类型参数（值类型）
func fun3(c [3]int) {
	fmt.Println("传入的参数c的值:", c)
	c[0] = 100
	c[1] = 200
	c[2] = 300
	fmt.Println("fun3()函数中修改了c的值，c的值:", c)
}

// fun4 传递指针类型参数，数组的指针（引用类型）
func fun4(d *[3]int) {
	fmt.Println("传入的参数d的值:", *d)
	(*d)[0] = 4
	(*d)[1] = 5
	(*d)[2] = 6
	fmt.Println("fun4()函数中修改了d的值，d的值:", *d)
}
