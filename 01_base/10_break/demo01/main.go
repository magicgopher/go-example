package main

import (
	"fmt"
	"time"
)

func main() {
	// break关键字的作用
	// 跳出循环： 当 break 语句出现在循环体（for 循环）中时，它会立即终止当前循环的执行，并跳出循环体，继续执行循环之后的代码。
	// 跳出 switch 语句： 当 break 语句出现在 switch 语句中时，它会立即终止当前 switch 语句的执行，并跳出 switch 语句，继续执行 switch 语句之后的代码。

	// 跳出for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // 当 i 等于 5 时，跳出循环
		}
		fmt.Println(i)
	}
	fmt.Println("循环结束")

	// 跳出switch语句
	switch day := time.Now().Weekday(); day {
	case time.Saturday:
		fmt.Println("今天是星期六")
		break // 跳出 switch 语句
	case time.Sunday:
		fmt.Println("今天是星期日")
		break // 跳出 switch 语句
	default:
		fmt.Println("今天是工作日")
	}
}
