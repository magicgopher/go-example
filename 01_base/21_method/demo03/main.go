package main

import "fmt"

func main() {
	/*
		值接收者和指针接收者

		实际开发过程如何选择接收者类型？是采用值类型还是指针类型？
		默认倾向于使用指针类型接收者
			原因：
				1.避免复制大型结构体的开销，提高性能。
				2.允许方法修改接收者的状态，灵活性更高。
				3.Go 标准库中许多类型（如 io.Reader, os.File）的方法都使用指针接收者，保持一致性是一个好习惯。
			场景：
				1.结构体较大或复杂。
				2.方法需要修改结构体字段。
				3.类型设计为面向对象风格（如实现接口时需要修改状态）。

		使用值类型接收者的情况
			原因：
				1.结构体较小，复制成本可以忽略（如 struct { x, y int }）。
				2.方法是纯函数（只读操作），不需要修改接收者。
				3.强调不可变性，避免意外修改原始对象。
			场景：
				1.小型、不可变的数据结构（如 time.Time）。
				2.方法仅用于计算或返回新值，不涉及状态变更。
		一致性优先
			如果一个类型的大多数方法需要修改状态（使用指针接收者），那么即使某些方法是只读的，也建议统一使用指针接收者。这种一致性可以减少开发者在使用类型时的困惑。
			反例：如果一个类型中混杂了值接收者和指针接收者，可能导致调用时的行为不直观。
	*/
	p := Person{"张三"}
	res1 := p.SetName1("李四")
	fmt.Println(res1)   // 张三
	fmt.Println(p.Name) // 李四

	fmt.Println("===============")

	p1 := &Person{"张三"}
	res2 := p1.SetName2("李四")
	fmt.Println(res2)    // 李四
	fmt.Println(p1.Name) // 李四
}

type Person struct {
	Name string
}

// 值接收者
func (p Person) SetName1(name string) string {
	p.Name = name
	return p.Name
}

// 指针接收者
func (p *Person) SetName2(name string) string {
	p.Name = name
	return p.Name
}
