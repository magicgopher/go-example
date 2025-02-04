package main

import (
	"fmt"
)

func main() {
	// 数据类型转换：Type Convert
	// go语言是静态语言，定义，赋值，运算必须类型一致
	// 语法格式：Type(Value)
	// 注意点：兼容类型可以转换，并非所有类型之间都可以相互转换。例如，你不能将一个字符串直接转换为一个整数。
	// 常数：在有需要的时候，自动转型
	// 变量：需要手动转型

	var a int8 = 123
	var b int64
	b = int64(a) // 将int8类型转为int64类型，也必须显示的转换
	fmt.Println(b)

	var c float64 = 3.14159
	d := int(c) // 将float64类型转为int类型会损失精度
	fmt.Println(d)

	// 必须是兼容类型才可以转换，不可以将一个布尔类型转为字符串类型
	//var f = true
	//str := string(f) // cannot convert f (variable of type bool) to type string
}
