package main

import "fmt"

func main() {
	/*
		string类型：字符串类型，本质就是[]byte，创建之后无法对其内容进行修改。

		Go中的字符串是Unicode兼容的，并且是UTF-8编码的。

		语法："" 或者 ``

		byte类型 ->unit8 类型 是兼容的，可以相互转换。
	*/

	str1 := "Hello World"
	str2 := `hello world`
	fmt.Println(str1)
	fmt.Println(str2)

	// 使用 len() 函数获取字符串的长度
	fmt.Println(len(str1))
	fmt.Println(len(str2))

	fmt.Println("==========")

	for i := 0; i < len(str1); i++ {
		fmt.Printf("i:%d, v:%v\n", i, str1[i])
	}

	fmt.Println("==========")

	// 使用 for range 遍历字符串的字符
	for i, v := range str2 {
		fmt.Printf("i:%d, v:%v\n", i, v)
	}

	fmt.Println("==========")

	slice1 := []byte{65, 66, 67, 68, 69}
	str3 := string(slice1) // 将[]byte切片转为string类型
	fmt.Println(str3)

	fmt.Println("==========")

	str4 := "Google"
	slice2 := []byte(str4)
	fmt.Println(slice2) // 将字符串转为[]byte切片

	// 根据索引下标获取字符串中的某个字符
	fmt.Println(str4[2])
	//str4[2] = 'O' // 尝试修改
	//fmt.Println(str4[2]) // cannot assign to str4[2]

	// 字符串拼接
	str5 := "Cloud"
	fmt.Println(str4 + str5)
}
