package main

import "fmt"

func main() {
	/*
		map数据结构属于复合类型
	*/

	map1 := make(map[int]string)
	map2 := make(map[string]float64)
	fmt.Printf("%T\n", map1)
	fmt.Printf("%T\n", map2)

	map3 := make(map[string]map[string]string)
	user1 := make(map[string]string)
	// 用户1
	user1["name"] = "张三"
	user1["age"] = "25"
	user1["sex"] = "男性"
	user1["address"] = "广东广州"
	// 职位1
	map3["golang开发工程师"] = user1
	// 用户2
	user2 := make(map[string]string)
	user2["name"] = "李四"
	user2["age"] = "23"
	user2["sex"] = "女性"
	user2["address"] = "广东广州"
	// 职位2
	map3["前端开发工程师"] = user2
	fmt.Println(map3)
}
