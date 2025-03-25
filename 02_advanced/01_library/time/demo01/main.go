package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("当前时间 (本地时区):", now)
	fmt.Println("当前时间 (UTC):", now.UTC())
}
