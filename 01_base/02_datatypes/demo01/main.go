package main

import "fmt"

// go基本类型有布尔型、数值型（整数型、浮点型、复数型）、字符串型、字节型、rune

func main() {
	// 数值类型（整数型）
	fmt.Println("数值类型-整数型-无符号位")
	// 无符号整数位
	var n1 uint = 100
	var n2 uint8 = 255
	var n3 uint16 = 65535
	var n4 uint32 = 4294967295
	var n5 uint64 = 18446744073709551615
	fmt.Println("uint类型变量n1的值是:", n1)
	fmt.Println("uint8类型变量n2的值是:", n2)
	fmt.Println("uint16类型变量n3的值是:", n3)
	fmt.Println("uint32类型变量n4的值是:", n4)
	fmt.Println("uint64类型变量n5的值是:", n5)
	fmt.Println("============")
	fmt.Println("数值类型-整数型-无符号位")
	// 有符号整数位
	var n6 int = -100
	var n7 int8 = -128
	var n8 int16 = -32768
	var n9 int32 = -2147483648
	var n10 int64 = -9223372036854775808
	fmt.Println("int类型变量n6的值是:", n6)
	fmt.Println("int8类型变量n7的值是:", n7)
	fmt.Println("int16类型变量n8的值是:", n8)
	fmt.Println("int32类型变量n9的值是:", n9)
	fmt.Println("int64类型变量n10的值是:", n10)
}
