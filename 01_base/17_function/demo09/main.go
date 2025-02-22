package main

import "fmt"

func main() {
	// 斐波那契数
	for i := 0; i <= 10; i++ {
		fmt.Printf("F(%d) = %d\n", i, fibonacci(i))
	}

	// 测试用例
	fmt.Println("getSum(5) =", getSum(5)) // 期望输出 15
	fmt.Println("getSum(3) =", getSum(3)) // 期望输出 6
	fmt.Println("getSum(1) =", getSum(1)) // 期望输出 1
	fmt.Println("getSum(0) =", getSum(0)) // 期望输出 0
}

// fibonacci 斐波那契数
// 时间复杂度为 O(1)
func fibonacci(num int) int {
	if num <= 0 {
		return 0
	}
	if num == 1 {
		return 1
	}
	return fibonacci(num-1) + fibonacci(num-2)
}

// 求第n项的和
// 时间复杂度为 O(1)
func getSum(num int) int {
	// 边界条件
	if num <= 0 {
		return 0 // 对于无效输入返回 0
	}
	if num == 1 {
		return 1 // 递归终止条件
	}
	// 递归公式：getSum(n) = getSum(n-1) + n
	return getSum(num-1) + num
}
