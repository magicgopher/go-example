package main

import "fmt"

func main() {
	// if...else...语句表示当条件表达式结果为true执行if {} 大括号内的语句体， 条件表达式结果为false执行else {} 大括号内的语句体
	/*
		if...else... 语法格式
			if 初始化语句; 条件表达式 {
				语句体;
			} else {
				语句体;
			}
	*/

	// 判断成绩是否及格
	if score := 59; score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}
	fmt.Println("main over")
}
