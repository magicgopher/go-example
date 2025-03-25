package main

import (
	"fmt"
	"time"
)

func main() {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, "2025-03-25 14:33:33")
	if err != nil {
		fmt.Println("解析错误:", err)
		return
	}
	fmt.Println("解析时间:", t) // 输出：2025-03-25 14:33:33 +0000 UTC
}
