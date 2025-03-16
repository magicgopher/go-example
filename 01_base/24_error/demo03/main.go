package main

import (
	"errors"
	"fmt"
)

func main() {
	// 创建一个error类型
	err := errors.New("这是一个错误.")
	fmt.Printf("%T\n", err) // *errors.errorString
	fmt.Println(err)

	fmt.Println("===============")

	err2 := fmt.Errorf("错误的信息码:%d", 100)
	fmt.Printf("%T\n", err2) // *errors.errorString
	fmt.Println(err2)

	fmt.Println("===============")

	err3 := checkAge(-10)
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	fmt.Println("main over")
}

// 检查年龄
func checkAge(age int) error {
	if age < 0 {
		return errors.New("年龄不能为负")
	}
	fmt.Println("年龄:", age)
	return nil
}
