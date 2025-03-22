package main

import (
	"errors"
	"fmt"
)

// 定义一个自定义错误类型
var ErrNotFound = errors.New("item not found")

func findItem(id int) error {
	// 模拟查找失败的情况
	return fmt.Errorf("findItem: %w", ErrNotFound) // 使用 %w 包装错误
}

func main() {
	err := findItem(42)
	if err != nil {
		// 使用 errors.Is 检查错误
		if errors.Is(err, ErrNotFound) {
			fmt.Println("The item was not found.")
		} else {
			fmt.Println("An unexpected error occurred:", err)
		}
	}
}
