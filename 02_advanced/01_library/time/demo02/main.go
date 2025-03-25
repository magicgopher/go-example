package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2025, time.March, 25, 10, 30, 0, 0, time.UTC)
	fmt.Println("指定时间:", t)
}
