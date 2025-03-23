package main

import "fmt"

func main() {
	// fmt包的错误格式化函数
	err := fmt.Errorf("用户 %s 不存在", "Alice") // 返回 error: 用户 Alice 不存在
	if err != nil {
		fmt.Println(err)
		return
	}
}
