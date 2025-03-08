package main

import "fmt"

func main() {
	/*
		通过匿名结构体的方式模拟继承的操作
		子结构体实例调用父结构体的方法

		如果存在方法的重写，子类对象访问重写的方法（也就是子结构体重写的方法）
		如果不存在方法的重写，子类对象访问父类的方法
	*/
	p1 := person{"张三", 20}
	fmt.Println(p1) // {张三 20}
	p1.eat()        // 父类实例访问父类的方法

	// student
	s1 := student{person{"李四", 16}, "001"}
	fmt.Println(s1) // {李四 16 001}
	s1.study()      // 子结构体实例自己的方法
	s1.person.eat() // 通过子结构体实例访问父类的方法
	s1.eat()        // 子结构体访问了重写父结构体的方法
}

// person 人
type person struct {
	name string // 姓名
	age  uint   // 年龄
}

// student 学生
type student struct {
	person        // 继承person的字段
	sid    string // 学号
}

// 方法
// eat
func (p person) eat() {
	fmt.Println("person eat")
}

// study
func (s student) study() {
	fmt.Println("student study")
}

// eat 重写了父结构体的方法
func (s student) eat() {
	fmt.Println("student eat")
}
