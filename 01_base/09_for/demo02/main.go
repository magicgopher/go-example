package main

import "fmt"

func main() {
	// 练习1：打印58-23数字
	for i := 58; i >= 23; i-- {
		fmt.Println(i) // 输出 58 至 23 的数值
	}
	fmt.Println("------------------")

	// 练习2：求1-100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum = sum + i
	}
	fmt.Println("1-100的和:", sum)
	fmt.Println("------------------")

	// 练习3：打印1-100内，能够被3整除，但是不能被5整除的数字，统计被打印的数字的个数，每行打印5个
	count := 0
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 != 0 {
			fmt.Printf("i:%d\t", i)
			count++
			if count%5 == 0 {
				fmt.Println()
			}
		}
	}
}
