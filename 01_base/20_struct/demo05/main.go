package main

import "fmt"

func main() {
	/*
		匿名结构体：是指没有显式定义的结构体，直接在代码种声明并使用，通常都是用于一次性或者临时的数据结构
		定义格式：
			变量名称 := struct {
						定义字段Field
					}{
						字段进行赋值
					}

		匿名字段：是指结构体定义中嵌入一个类型（通常式其他结构图或者接口类型）
				但不显式指定字段名称，这时，嵌入的类型本身即是字段类型，也是字段名称。
	*/

	// 结构体实例
	s1 := Student{Person{"张三", 17}, "001"}
	s2 := Student{Person{"李四", 18}, "002"}
	s3 := Student{Person{"王五", 19}, "003"}
	//fmt.Println(s1, s2, s3)

	// 创建切片保存学生实例
	students := make([]Student, 0, 10)
	students = append(students, s1, s2, s3)

	// 创建一个老师的匿名结构体
	t1 := struct {
		name     string
		age      uint
		students []Student
	}{
		"王老师",
		18,
		students,
	}
	fmt.Println(t1)
}

// Student 学生
type Student struct {
	Person        // 匿名字段：嵌入 Person 结构体
	sid    string // 学号
}

// Person 人
type Person struct {
	name string // 姓名
	age  uint   // 年龄
}
