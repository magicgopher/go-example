package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("开始:", time.Now())
	time.Sleep(2 * time.Second) // 主goroutine睡眠2秒
	fmt.Println("结束:", time.Now())
}
