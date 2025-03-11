package main

import "fmt"

func main() {
	/*
		接口嵌套
	*/

	u := User{}
	u.TestA()
	u.TestB()
	u.TestC()

	fmt.Println("==============")

	var a A = User{} // 这里可以调用接口A的方法
	a.TestA()

	fmt.Println("==============")

	var b B = User{} // 这里可以调用接口B的方法
	b.TestB()

	fmt.Println("==============")

	var c C = User{} // 这里可以调用接口C、接口A、接口B的方法
	c.TestA()
	c.TestB()
	c.TestC()
}

// A 接口A
type A interface {
	TestA() // 接口A的方法
}

// B 接口B
type B interface {
	TestB() // 接口B的方法
}

// C 接口C
type C interface {
	A       // 嵌套接口A
	B       // 嵌套接口B
	TestC() // 接口C的方法
}

// User 用户
type User struct {
	// 如果这个结构体想实现C接口的方法，那么还需要实现接口A和接口B的方法
}

func (u User) TestA() {
	fmt.Println("User TestA")
}

func (u User) TestB() {
	fmt.Println("User TestB")
}

func (u User) TestC() {
	fmt.Println("User TestC")
}
