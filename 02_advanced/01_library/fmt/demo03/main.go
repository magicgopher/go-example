package main

import "fmt"

func main() {
	// fmt包的输入函数

	// Scan
	//fun1()

	// Scanln
	//fun2()

	// Scanf
	fun3()
}

func fun1() {
	var name string
	var age int
	fmt.Println("请输出姓名和年龄，例子： Alice 25")
	fmt.Scan(&name, &age)      // 输入: Alice 25
	fmt.Println("Name:", name) // 输出: Name: Alice
	fmt.Println("Age:", age)   // 输出: Age: 25
}

func fun2() {
	var name string
	var age int
	fmt.Println("请输出姓名和年龄，例子：Alice 25\n")
	fmt.Scanln(&name, &age)    // 输入: Alice 25\n
	fmt.Println("Name:", name) // 输出: Name: Alice
	fmt.Println("Age:", age)   // 输出: Age: 25
}

func fun3() {
	var name string
	var age int
	fmt.Println("请输出姓名和年龄，例子：Name: Alice, Age: 25")
	fmt.Scanf("Name: %s, Age: %d", &name, &age) // 输入: Name: Alice, Age: 25
	fmt.Println("Name:", name)                  // 输出: Name: Alice
	fmt.Println("Age:", age)                    // 输出: Age: 25
}
