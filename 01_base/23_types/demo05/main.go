package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		扩展 time.Time，添加检查是否为周末的方法。
	*/
	now := MyTime(time.Now())
	fmt.Printf("Today is weekend: %t\n", now.IsWeekend())
}

// 类型定义：MyTime 基于 time.Time
type MyTime time.Time

func (mt MyTime) IsWeekend() bool {
	t := time.Time(mt)
	weekday := t.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}
