package main

import "fmt"

func main() {
	/*
		空接口没有定义任何方法，因此任何类型都隐式实现了空接口。空接口常用于表示任意类型
	*/
	var a1 Any = User{"Martin", 18}
	var a2 Any = Cat{"黑色"}
	var a3 Any = 100
	var a4 = "Golang"
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

	test1(a1)
	test1(a2)
	test1(3.14)
	test1("李四")

	test2(a3)
	test2(1000)

	//map，key字符串，value任意类型
	map1 := make(map[string]interface{})
	map1["name"] = "Allan"
	map1["age"] = 30
	map1["friend"] = User{"Jerry", 18}
	fmt.Println(map1)

	//切片，存储任意类型的数据
	slice1 := make([]interface{}, 0, 10)
	slice1 = append(slice1, a1, a2, a3, a4, 100, "abc", true, 9.88, 'd', User{"Tom", 18}, Cat{"白色"})
	fmt.Println(slice1)

	test3(slice1)
}

// Any 表示空接口，可以存储任何类型
type Any interface {
}

type User struct {
	Name string // 姓名
	Age  uint   // 年龄
}

type Cat struct {
	Color string // 颜色
}

func test3(slice2 []interface{}) {
	for i := 0; i < len(slice2); i++ {
		fmt.Printf("第%d个数据：%v\n", i+1, slice2[i])
	}
}

// 接口A是空接口，理解为代表了任意类型
func test1(a Any) {
	fmt.Println(a)
}

func test2(a interface{}) {
	fmt.Println("--->", a)
}
