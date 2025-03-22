package main

import "fmt"

func main() {
	// 第二种方式创建 error 通过 fmt 包的 Errorf() 函数
	// res2, err := sum2(11, 22)
	res2, err := sum2(-11, -22)
	if err != nil {
		panic(err)
	}
	fmt.Println(res2)
}

// sum2
func sum2(a, b int) (int, error) {
	if a < 0 && b < 0 {
		return 0, fmt.Errorf("参数必须大于0")
	}
	return a + b, nil
}
