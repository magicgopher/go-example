package main

import "fmt"

func main() {
	// 数值型（复数型）
	var c1 complex64 = 3.14159265358979 + 2.71828182845905i
	var c2 complex128 = 3.14159265358979323846264338327950288 + 2.71828182845904523536028747135266249i

	fmt.Println(c1)
	fmt.Println(c2)
}
