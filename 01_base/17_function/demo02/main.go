package main

import (
	"fmt"
	"strconv"
)

// Add 将两个整数相加
// 命名依据：大写开头（尽管在 main 包中无导出需求，仍保持清晰），动词表示操作
func Add(a, b int) int {
	return a + b
}

// subtract 减去两个整数，模拟私有功能
// 命名依据：小写开头（包内使用），动词表示操作
func subtract(a, b int) int {
	return a - b
}

// MultiplyAndFormat 乘法并格式化结果为字符串
// 命名依据：大写开头，驼峰命名，动词+描述性名词
func MultiplyAndFormat(a, b int) string {
	result := a * b
	return fmt.Sprintf("Result: %d", result)
}

// parseNumber 将字符串解析为整数，辅助函数
// 命名依据：小写开头，动词+名词，清晰表达功能
func parseNumber(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %v", err)
	}
	return num, nil
}

// CalculateSumAndDiff 计算和与差
// 命名依据：大写开头，驼峰命名，动词+具体功能描述
func CalculateSumAndDiff(a, b int) (sum int, diff int) {
	sum = Add(a, b)       // 调用 Add 函数
	diff = subtract(a, b) // 调用 subtract 函数
	return
}

func main() {
	/*
		函数的命名规范
			1.见名知意，避免冗长或模糊，遵循“短小精悍”的原则。
			2.函数名大小写控制可见性
				大写开头（如 Add）：函数是公开的，可以被包外访问。
				小写开头（如 add）：函数是私有的，仅包内可见。
			3.Go 函数名采用驼峰命名法（CamelCase），即首字母可以大写或小写，后续单词的首字母大写，不使用下划线 _ 分隔单词。
			4.函数名应尽量避免缩写，除非缩写是领域内广泛接受的术语（如 ID、HTTP、JSON）。
	*/
	// 示例 1：调用 Add
	sum := Add(3, 5)
	fmt.Println("Sum:", sum) // 输出: Sum: 8

	// 示例 2：调用 MultiplyAndFormat
	result := MultiplyAndFormat(4, 2)
	fmt.Println(result) // 输出: Result: 8

	// 示例 3：调用 CalculateSumAndDiff
	s, d := CalculateSumAndDiff(10, 7)
	fmt.Println("Sum:", s, "Diff:", d) // 输出: Sum: 17 Diff: 3

	// 示例 4：调用 parseNumber
	num1, _ := parseNumber("42")
	fmt.Println("Parsed number:", num1) // 输出: Parsed number: 42
}
