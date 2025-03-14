package main

import (
	"errors"
	"fmt"
)

func main() {
	/*
		错误处理：
	*/

	// 除法运算
	//res := sub(10, 5)
	//res := sub(10, 0) // panic: runtime error: integer divide by zero
	//fmt.Println(res)

	res, err := sub(10, 5)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", res)
}

// 除法运算
//func sub(a, b int) int {
//	result := a / b
//	return result
//}

// 除法运算
func sub(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	result := a / b
	return result, nil
}
