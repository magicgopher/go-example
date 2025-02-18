package main

import "fmt"

func main() {
	/*
		map和slice结合使用案例
			1.使用map用于存储个人信息
			2.每个map存储一个人的信息
			3.将这些map存入到slice中
			4.打印遍历这些个人信息
	*/

	user1 := map[string]string{
		"name":    "张三",
		"age":     "18",
		"sex":     "男性",
		"address": "广东广州",
	}
	user2 := map[string]string{
		"name":    "李四",
		"age":     "20",
		"sex":     "女性",
		"address": "浙江杭州",
	}
	user3 := map[string]string{
		"name":    "王五",
		"age":     "16",
		"sex":     "男性",
		"address": "四川成都",
	}
	// 将map存入slice中
	users := make([]map[string]string, 0, 5)
	users = append(users, user1)
	users = append(users, user2)
	users = append(users, user3)

	fmt.Println("users:", users)
}
