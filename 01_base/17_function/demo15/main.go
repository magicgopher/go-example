package main

import "fmt"

func main() {
	/*
		go语言支持函数式编程：
				支持将一个函数作为另一个函数的参数，
				也支持将一个函数作为另一个函数的返回值。

			闭包(closure)：
				一个外层函数中，有内层函数，该内层函数中，会操作外层函数的局部变量(外层函数中的参数，或者外层函数中直接定义的变量)，并且该外层函数的返回值就是这个内层函数。

				这个内层函数和外层函数的局部变量，统称为闭包结构。


				局部变量的生命周期会发生改变，正常的局部变量随着函数调用而创建，随着函数的结束而销毁。
				但是闭包结构中的外层函数的局部变量并不会随着外层函数的结束而销毁，因为内层函数还要继续使用。
	*/
	// 调用 outerFunc，获取返回的函数
	greet := outerFunc("Hello")

	// 使用返回的函数
	result := greet("Alice")
	fmt.Println(result) // 输出: Hello Alice

	result2 := greet("Bob")
	fmt.Println(result2) // 输出: Hello Bob

	logInfo := createLogger("INFO")
	logError := createLogger("ERROR")

	logInfo("User logged in")  // 输出: [INFO] User logged in
	logError("Server crashed") // 输出: [ERROR] Server crashed
}

// outerFunc 返回一个函数类型
func outerFunc(prefix string) func(string) string {
	return func(name string) string {
		return prefix + " " + name
	}
}

// createLogger
func createLogger(level string) func(string) {
	return func(message string) {
		fmt.Printf("[%s] %s\n", level, message)
	}
}
