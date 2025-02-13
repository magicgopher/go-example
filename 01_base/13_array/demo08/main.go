package main

import "fmt"

func main() {
	// 多维数组：本质是存储数组的数组，可以把它们看作是嵌套的数组结构。最常见的多维数组是二维数组，类似于表格或矩阵。
	// 语法格式：var array_name [size1][size2]...[sizeN] type
	// array_name：是数组的名称
	// size1, size2, ..., sizeN 分别是各个维度的长度。
	// type：是数组元素的类型。

	// 定义一个二维数组
	// [2]表示：数组的第一维（行）的长度为 2。
	// [3]表示：表示数组的第二维（列）的长度为 3。
	var arr1 = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	// 访问数组元素
	fmt.Println(arr1[0][0]) // 输出: 1
	fmt.Println(arr1[1][2]) // 输出: 6
	fmt.Println("==========")

	// 遍历数组
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", arr1[i][j])
		}
		fmt.Println()
	}
	fmt.Println("==========")

	// 定义一个三维数组
	// [2]表示：数组的第一维（可以理解为“层”）的长度为 2。
	// [3]表示：数组的第二维（可以理解为“行”）的长度为 3。
	// [4]表示：数组的第三维（可以理解为“列”）的长度为 4。
	arr2 := [2][3][4]int{
		{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		},
		{
			{13, 14, 15, 16},
			{17, 18, 19, 20},
			{21, 22, 23, 24},
		},
	}

	// 访问数组元素
	fmt.Println(arr2[0][1][2]) // 输出: 7
	fmt.Println(arr2[1][2][3]) // 输出: 24
	fmt.Println("==========")

	// 遍历数组
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 4; k++ {
				fmt.Printf("%d ", arr2[i][j][k])
			}
			fmt.Println()
		}
		fmt.Println()
	}
	fmt.Println("main over")
}
