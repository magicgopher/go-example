package main

import "fmt"

func main() {
	/*
		函数的返回值：一个函数的执行结果，返回给函数的调用处。执行结果就叫做函数的返回值。

		return关键字：一个函数有返回值，可以使用return关键字，将函数处理的结果，也就是返回值返回给调用方。
			函数返回的结果，必须和函数定义的一致：类型，个数，顺序。
			1.将函数的结果返回给调用处
			2.同时结束了该函数的执行

		空白标识符，专门用于舍弃数据：_
	*/

	// 调用函数
	_, area := CalculateRect(5.2, 3) // 面积
	fmt.Printf("面积:%0.3f\n", area)

	perimeter, _ := CalculateRect(5.2, 3) // 周长
	fmt.Printf("周长:%0.3f\n", perimeter)

	perimeter1, area2 := CalculateRect(5.2, 3) // 两个返回值
	fmt.Printf("周长:%0.3f, 面积:%0.3f\n", perimeter1, area2)
}

// CalculateRect 计算长方形的周长和面积
func CalculateRect(length, width float64) (perimeter float64, area float64) {
	// 检查输入是否有效（长度和宽度必须大于0）
	if length <= 0 || width <= 0 {
		return 0, 0 // 返回零值表示错误
	}

	// 周长 = 2 * (长 + 宽)
	perimeter = 2 * (length + width)

	// 面积 = 长 * 宽
	area = length * width

	return perimeter, area
}
