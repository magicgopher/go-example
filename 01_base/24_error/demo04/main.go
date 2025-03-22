package main

import (
	"errors"
	"fmt"
)

func main() {
	// 第一种方式创建 error 通过 errors 包的 New() 函数
	// res1, err := sum1(10, 20)
	res1, err := sum1(-10, -20)
	if err != nil {
		panic(err)
	}
	fmt.Println(res1)
}

// sum1
func sum1(a, b int) (int, error) {
	if a < 0 && b < 0 {
		return 0, errors.New("参数必须大于0")
	}
	return a + b, nil
}
