package main

import "fmt"

// b := 100 // 简短定义方式，不能定义全局变量
var b = 100

func main() {
	// 变量需要注意的事项
	// 1.变量必须先声明才能使用
	// 2.go是强类型语言，要求变量的类型和赋值的类型必须一致
	// 3.变量名不能冲突（在同一个作用域内不能冲突）
	// 4.简短定义方式，不能定义全局变量
	// 5.变量的零值。也叫默认值
	// 6.变量定义了就要使用，否则无法通过编译

	// 变量必须先声明才能使用
	//fmt.Println(b) // undefined: b

	// go是强类型语言，要求变量的类型和赋值的类型必须一致
	var a int
	// a = "haha" // cannot use "haha" (untyped string constant) as int value in assignment
	a = 100
	fmt.Println(a)

	//var a string // 在同一个作用域【这里的作用域名是指{}】重复声明变量a
	//fmt.Println(a)

	// 打印输出全局变量b
	fmt.Println(b)

	// 变量有默认值
	var n1 int // 0
	fmt.Println(n1)
	var n2 string // ""空字符串
	fmt.Println(n2)
	var n3 float64 // 0.0
	fmt.Println(n3)
	var n4 bool // false
	fmt.Println(n4)
	var n5 rune // 0
	fmt.Println(n5)

	//var d byte = 10 // 定义的变量未使用会编译不通过：declared and not used: d
}
