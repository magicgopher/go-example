package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		strings包提供了一系列用于操作字符串的函数
	*/

	// strings.Contains()：判断字符串 s 是否包含子串 str
	str1 := "Google Cloud"
	fmt.Println(strings.Contains(str1, "Google")) // 结果：true

	fmt.Println("==========")

	// strings.ContainsAny()：检查给定的字符串是否包含任何指定字符集中的字符
	str2 := "Hello, World"
	fmt.Println(strings.ContainsAny(str2, "xyz")) // 返回 false，因为 "Hello, World" 中没有包含 'x', 'y' 或 'z'
	fmt.Println(strings.ContainsAny(str2, "HW"))  // 返回 true，因为 "Hello, World" 中包含 'H' 或 'W'

	fmt.Println("==========")

	// strings.Count()：统计字符串 s 中子串 str 出现的次数
	str3 := "Microsoft AWS"
	fmt.Println(strings.Count(str3, "o")) // 结果：2

	fmt.Println("==========")

	// strings.Index()：找出子串在字符串中第一次出现的位置，如果不存在返回-1，区分大小写
	str4 := "google"
	fmt.Println(strings.Index(str4, "L"))     // 结果：-1
	fmt.Println(strings.Index(str4, "l"))     // 结果：4
	fmt.Println(strings.Index(str4, "hello")) // 结果：-1

	fmt.Println("==========")

	// strings.HasPrefix()：检查字符串是否以特定前缀开始。
	str5 := `Gopher`
	fmt.Println(strings.HasPrefix(str5, "go")) // 结果：false
	fmt.Println(strings.HasPrefix(str5, "Go")) // 结果：true

	fmt.Println("==========")

	// strings.HasSuffix()：检查字符串是否以特定后缀结束。
	str6 := `golang并发学习笔记.txt`
	fmt.Println(strings.HasSuffix(str6, ".txt")) // 结果：true
	fmt.Println(strings.HasSuffix(str6, ".md"))  // 结果：false

	fmt.Println("==========")

	// strings.Join()：用指定的连接符将字符串切片连接起来。
	slice1 := []string{"hello", "world", "golang"}
	fmt.Println(strings.Join(slice1, " ")) // hello world golang
	fmt.Println(strings.Join(slice1, "+")) // hello+world+golang

	fmt.Println("==========")

	// strings.Repeat()：用于将一个字符串重复指定的次数，并返回一个新的字符串。
	fmt.Println(strings.Repeat("c", 3))     // ccc
	fmt.Println(strings.Repeat("hello", 2)) // hellohello

	fmt.Println("==========")

	// strings.Replace()：用于在字符串中替换指定的子串。它可以替换所有出现的子串，也可以只替换指定数量的子串。
	str7 := "hello golang hello java hello python"
	fmt.Println(strings.Replace(str7, "hello", "hi", -1)) // -1 表示替换无数次
	fmt.Println(strings.Replace(str7, "hello", "hi", 1))  // 替换一次
	fmt.Println(strings.Replace(str7, "hello", "hi", 2))  // 替换两次

	fmt.Println("==========")

	// strings.Split()：函数用于将一个字符串分割成多个子字符串，并返回一个字符串切片。
	str8 := "apple,banana,orange"
	fmt.Println(strings.Split(str8, ",")) // [apple banana orange]
	str9 := "a b c"
	fmt.Println(strings.Split(str9, " ")) // [a b c]
	str10 := "hello"
	fmt.Println(strings.Split(str10, "")) // [h e l l o]

	fmt.Println("==========")

	// strings.Trim()：函数用于删除字符串开头和结尾指定字符集合中的字符。
	str11 := "   hello world!   "
	fmt.Println(strings.Trim(str11, " ")) // 输出：hello world!
	str12 := "!@#hello world!@#"
	fmt.Println(strings.Trim(str12, "!@#")) // 输出：hello world
	str13 := "/home/user/"
	fmt.Println(strings.Trim(str13, "/")) // 输出：home/user
}
