package main

import "fmt"

func main() {
	// 通过 [...] 替代数组的长度，这里会根据元素的个数来推断出数组的长度
	var array = [...]string{"Hello", "World", "你好", "世界"}
	fmt.Println(array)
	fmt.Println("长度:", len(array))
	fmt.Println("容量:", cap(array))
}
