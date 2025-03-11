package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		定义中间件函数类型，实现字符串处理链。
	*/

	upper := func(s string) string { return strings.ToUpper(s) }
	prefix := func(s string) string { return "prefix-" + s }
	result := applyMiddleware("hello", upper, prefix)
	fmt.Println(result) // 输出: prefix-HELLO
}

// 函数类型：Middleware 表示字符串处理函数
type Middleware func(string) string

func applyMiddleware(input string, mws ...Middleware) string {
	result := input
	for _, mw := range mws {
		result = mw(result)
	}
	return result
}
