package main

import "fmt"

func main() {
	// continue 关键字通常与条件语句（如 if 语句）一起使用，用于在满足特定条件时跳过当前循环迭代。
	// continue 关键字注意事项
	// 1.continue 关键字只能用于循环语句（for 循环）。
	// 2.当 continue 语句出现在嵌套循环中时，它只会跳过当前循环迭代的剩余语句，而不会影响外层循环的执行。
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // 如果 i 是偶数，则跳过本次循环
		}
		fmt.Println(i)
	}
}
