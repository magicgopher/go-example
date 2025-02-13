package main

import "fmt"

func main() {
	// 创建三个int类型的数组
	a1 := [3]int{11, 22, 33}
	a2 := [3]int{11, 22, 33}
	a3 := [3]int{67, 19, 49}
	// 使用 == 可以比较两个数组的元素是否相同
	fmt.Printf("a1 == a2 结果:%t\n", a1 == a2)
	fmt.Printf("a1 == a3 结果:%t\n", a1 == a3)

	// 注意事项：必须是两个数组类型相同才能进行比较。【数据类型相同：长度相同、数据类型相同】
	a4 := [...]int{1, 2, 3, 4}
	//fmt.Printf("a1 == a4 结果:%t\n", a1 == a4) // 编译不通过
	a5 := [...]int{1, 2, 3, 4}
	fmt.Printf("a4 == a5 结果:%t\n", a4 == a5)
	a6 := [...]int{7, 8, 9, 10}
	fmt.Printf("a4 == a6 结果:%t\n", a4 == a6)
}
