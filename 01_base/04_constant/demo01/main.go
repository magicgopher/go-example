package main

import "fmt"

func main() {
	// 常量：计算机的一小块内存，用于存储数据，在程序运行的过程中，数据是不可变的。
	// 常量的概念和变量有点类似，只是常量不可变，而变量可变。

	// 定义一个常量的语法格式：
	// 显式类型定义：const 常量名称 数据类型 = 赋值的数据
	// 隐式类型定义：const 常量名称 = 赋值的数据

	// 1.定义常量
	const URL string = "https://github.com/"
	const PI = 3.14
	fmt.Println(URL)
	fmt.Println(PI)

	// 2.尝试修改定义的常量
	//URL = "https://google.com/" // cannot assign to URL (neither addressable nor a map index expression)

	// 3.定义一组常量
	const C1, C2, C3 = 100, 5.55, "golang"

	// 4.使用()定义一组常量，如果某个常量没有初始值，默认和上一行一致
	const (
		a = 100
		b
		c = "hello"
		d
		e = 5.55
		f = true
	)
	fmt.Printf("%T,%d\n", a, a)
	fmt.Printf("%T,%d\n", b, b)
	fmt.Printf("%T,%s\n", c, c)
	fmt.Printf("%T,%s\n", d, d)
	fmt.Printf("%T,%0.2f\n", e, e)
	fmt.Printf("%T,%t\n", f, f)

	// 5.可以使用定义一组常量组作为枚举类型。一组相关数值的数据
	const (
		SPRING = 0
		SUMMER = 1
		AUTUMN = 2
		WINTER = 3
	)
}
