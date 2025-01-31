package main

import "fmt"

func main() {
	// iota：特殊的常量，可以被编译器修改的常量
	// 每当定义一个const，iota的初始值为0
	// 每当定义一个常量，就会自动累加1
	// 直到下一个const出现，清零

	// 定义一个iota
	const (
		A = "gopher" // iota = 0
		B = iota     // iota = 1
		C = 20       // iota = 2
		D = iota     // iota = 3
		E = false    // iota = 4
		F = iota     // iota = 5
	)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)

	const (
		G = iota // 这里有是一个新的const所以iota又从0开始
		H
		I
	)
	fmt.Println(G, H, I)

	// 枚举 iota
	const (
		MALE   = iota // 0
		FEMALE        // 1
		UNKNOW        // 2
	)
	fmt.Println(MALE, FEMALE, UNKNOW)
}
