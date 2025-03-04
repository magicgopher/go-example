package main

import "fmt"

func main() {
	/*
		多级指针：
			一级指针（*T）：指向某个类型 T 的变量。
			二级指针（**T）：指向一级指针的指针。
			三级指针（***T）：指向二级指针的指针，以此类推。

		Go 语言支持任意级别的指针，但实际开发中通常不会使用过多的层级（比如超过三级），因为代码会变得复杂且难以维护。
	*/

	// 一级
	one()

	fmt.Println("===============")

	// 二级
	two()

	fmt.Println("===============")

	// 三级
	three()
}

func one() {
	var a int = 10  // 定义一个整型变量
	var p *int = &a // 一级指针，指向 a 的地址

	fmt.Println("a 的值:", a)    // 输出: 10
	fmt.Println("p 指向的值:", *p) // 输出: 10（通过解引用获取 a 的值）
	fmt.Println("a 的地址:", &a)  // 输出: a 的内存地址
	fmt.Println("p 的值:", p)    // 输出: a 的内存地址
}

func two() {
	var a int = 10    // 定义一个整型变量
	var p *int = &a   // 一级指针，指向 a
	var pp **int = &p // 二级指针，指向 p

	fmt.Println("a 的值:", a)       // 输出: 10
	fmt.Println("p 指向的值:", *p)    // 输出: 10
	fmt.Println("pp 指向的值:", **pp) // 输出: 10（两次解引用）
	fmt.Println("p 的地址:", &p)     // 输出: p 的内存地址
	fmt.Println("pp 的值:", pp)     // 输出: p 的内存地址
}

func three() {
	var a int = 10       // 定义一个整型变量
	var p *int = &a      // 一级指针，指向 a
	var pp **int = &p    // 二级指针，指向 p
	var ppp ***int = &pp // 三级指针，指向 pp

	fmt.Println("a 的值:", a)          // 输出: 10
	fmt.Println("p 指向的值:", *p)       // 输出: 10
	fmt.Println("pp 指向的值:", **pp)    // 输出: 10
	fmt.Println("ppp 指向的值:", ***ppp) // 输出: 10（三次解引用）
}
