package main

import (
	"errors"
	"fmt"
)

// 定义一个自定义错误类型
type MyError struct {
	Message string
}

func (e *MyError) Error() string {
	return e.Message
}

// 一个函数返回 MyError 类型的错误
func doSomething() error {
	return &MyError{"something went wrong"}
}

func main() {
	err := doSomething()
	if err != nil {
		// 使用 errors.As 检查错误并将其转换为 *MyError 类型
		var myErr *MyError
		if errors.As(err, &myErr) {
			fmt.Println("Caught a MyError:", myErr.Message)
		} else {
			fmt.Println("An unexpected error occurred:", err)
		}
	}
}
