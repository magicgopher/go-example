package main

import "fmt"

func main() {
	/*
		使用结构体嵌套模拟继承关系、聚合关系

		继承关系：
			type A struct {
				field
				...
			}

			type B struct {
				A // 结构体中的匿名字段
				...
			}

		聚合关系：
			type C struct{
				field
			}
	*/

	emp := Employee{
		Person: Person{
			Name: "David",
			Age:  40,
		},
		JobTitle: "Manager",
		HomeAddr: Address{
			City:    "Paris",
			Country: "France",
		},
	}

	// 继承相关
	fmt.Println(emp.Greet())       // 方法提升
	fmt.Println("Name:", emp.Name) // 字段提升

	// 聚合相关
	fmt.Println("City:", emp.HomeAddr.City)
}

// Person 模拟父类
type Person struct {
	Name string // 姓名
	Age  uint   // 年龄
	Sex  string // 性别
}

// Greet 迎接
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, I'm %s.", p.Name)
}

// Employee 员工
type Employee struct {
	Person   // 匿名字段，模拟继承
	JobTitle string
	HomeAddr Address // 命名字段，模拟聚合
}

// Address 地址
type Address struct {
	City    string // 城市
	Country string // 国家
}
