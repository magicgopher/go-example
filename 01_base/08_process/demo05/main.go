package main

import "fmt"

func main() {
	/*
		if嵌套语法格式：
			if condition1 {
				// condition1 为真时执行的代码块
				if condition2 {
					// condition1 和 condition2 都为真时执行的代码块
				} else {
					// condition1 为真，condition2 为假时执行的代码块
				}
			} else {
		    	// condition1 为假时执行的代码块
			}
	*/

	x := 10
	y := 5

	if x > y {
		fmt.Println("x 大于 y")
		if x > 5 {
			fmt.Println("x 大于 5")
		} else {
			fmt.Println("x 不大于 5")
		}
	} else {
		fmt.Println("x 不大于 y")
	}
}
