package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		strconv包：提供了一组函数，可用于处理各种类型的数据转换。
	*/

	// 字符串转为int
	atoi, err := strconv.Atoi("123")
	if err != nil {
		panic(err) // 错误处理
	}
	fmt.Printf("类型:%T, 值:%v\n", atoi, atoi)
	// int转字符串
	fi1 := strconv.FormatInt(456, 2) // 转为2进制字符串数值
	fmt.Printf("类型:%T, 值:%v\n", fi1, fi1)
	fi2 := strconv.FormatInt(456, 8) // 转为8进制字符串数值
	fmt.Printf("类型:%T, 值:%v\n", fi2, fi2)
	fi3 := strconv.FormatInt(456, 10) // 转为10进制字符串数值
	fmt.Printf("类型:%T, 值:%v\n", fi3, fi3)
	fi4 := strconv.FormatInt(456, 16) // 转为16进制字符串数值
	fmt.Printf("类型:%T, 值:%v\n", fi4, fi4)
	// 将int转字符串类型数值
	itoa := strconv.Itoa(123) // int转为10进制字符串数值
	fmt.Printf("类型:%T, 值:%v\n", itoa, itoa)

	fmt.Println("====================")

	// 字符串转浮点数
	f1, _ := strconv.ParseFloat("8.88", 64) // 字符串转float64类型
	fmt.Printf("类型:%T, 值:%v\n", f1, f1)
	f2, _ := strconv.ParseFloat("8.88", 32)
	fmt.Printf("类型:%T, 值:%v\n", f2, f2)

	f := 123.456789

	// 使用 'f' 格式，保留 2 位小数，float64类型
	fs1 := strconv.FormatFloat(f, 'f', 2, 64)
	fmt.Printf("类型:%T, 值:%v\n", fs1, fs1)
	// 使用 'e' 格式，科学计数法，保留 3 位小数 float32类型
	fs2 := strconv.FormatFloat(f, 'e', 3, 32)
	fmt.Printf("类型:%T, 值:%v\n", fs2, fs2)
	// 使用 'b' 格式，二进制指数 float64类型
	fs3 := strconv.FormatFloat(f, 'b', -1, 64)
	fmt.Printf("类型:%T, 值:%v\n", fs3, fs3)

	fmt.Println("====================")

	pb1, _ := strconv.ParseBool("t")
	fmt.Printf("类型:%T, 值:%v\n", pb1, pb1)
	pb2, _ := strconv.ParseBool("true")
	fmt.Printf("类型:%T, 值:%v\n", pb2, pb2)
	pb3, _ := strconv.ParseBool("f")
	fmt.Printf("类型:%T, 值:%v\n", pb3, pb3)
	pb4, _ := strconv.ParseBool("false")
	fmt.Printf("类型:%T, 值:%v\n", pb4, pb4)

	fb1 := strconv.FormatBool(true)
	fmt.Printf("类型:%T, 值:%v\n", fb1, fb1)
	fb2 := strconv.FormatBool(false)
	fmt.Printf("类型:%T, 值:%v\n", fb2, fb2)

	fmt.Println("====================")

	pu1, _ := strconv.ParseUint("12345", 10, 64) // 字符串转10进制uint64类型
	fmt.Printf("类型:%T, 值:%v\n", pu1, pu1)
	pu2, _ := strconv.ParseUint("0xFF", 0, 64)
	fmt.Printf("类型:%T, 值:%v\n", pu2, pu2)

	fu1 := strconv.FormatUint(65535, 2) // uint类型转二进制格式字符串
	fmt.Printf("类型:%T, 值:%v\n", fu1, fu1)
	fu2 := strconv.FormatUint(65535, 8) // uint类型转8进制格式字符串
	fmt.Printf("类型:%T, 值:%v\n", fu2, fu2)
	fu3 := strconv.FormatUint(65535, 10) // uint类型转10进制格式字符串
	fmt.Printf("类型:%T, 值:%v\n", fu3, fu3)
	fu4 := strconv.FormatUint(65535, 16) // uint类型转16进制格式字符串
	fmt.Printf("类型:%T, 值:%v\n", fu4, fu4)
}
