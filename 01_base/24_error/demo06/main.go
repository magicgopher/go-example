package main

import (
	"errors"
	"fmt"
)

func main() {
	// 错误传递
	// 在函数调用中传递错误是Go语言处理错误的常见方式。
	// 通过返回值返回错误，并在调用方进行检查。
	err := doSomething()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func doSomething() error {
	// 一些操作
	if true {
		return errors.New("发生了错误")
	}
	return nil
}
