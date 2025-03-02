package main

import "fmt"

func main() {
	// 空指针
	var p1 *string
	fmt.Printf("类型:%T\n", p1)
	fmt.Println(p1 == nil)

	// 通过*访问指针变量存储的地址存储的值
	var a = 123      // 声明并初始化一个整型变量 a，赋值为 123
	fmt.Println(a)   // 打印变量 a 的值，此时输出 123
	var p2 *int      // 声明一个整型指针变量 p2，初始值为 nil
	p2 = &a          // 将 p2 指向变量 a 的内存地址，&a 表示取 a 的地址
	fmt.Println(*p2) // 解引用 p2，打印 p2 指向的内存地址中的值，此时输出 123
	*p2 = 1234       // 通过指针 p2 修改 a 的值，将其赋值为 1234
	fmt.Println(a)   // 再次打印 a 的值，由于 p2 修改了 a 的值，此时输出 1234
}
