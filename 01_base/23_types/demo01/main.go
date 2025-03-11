package main

import "fmt"

func main() {
	/*
		type关键字：在 Go 中 type 关键字用于定义自定义类型。

		语法格式：
			1.类型定义：type 类型名 类型定义
			2.类型别名：type 类型名 = 类型定义
	*/
	var now Timestamp = 1698765432 // 假设这是一个时间戳
	logEvent(now)                  // 输出:"Event logged at: 1698765432"
}

type Timestamp int64 // 自定义类型

func logEvent(t Timestamp) {
	fmt.Printf("Event logged at: %d\n", t)
}
