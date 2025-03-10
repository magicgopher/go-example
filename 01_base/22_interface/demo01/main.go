package main

func main() {
	/*
		接口是一组方法的集合，可以用来定义一些行为的规范，是一种抽象的概念，是一种类型

		当某个类型为了这个接口中的所有方法提供了方法实现，那么这被称为实现接口。

		Go语言中接口和类型的实现关系是非嵌入式的。
	*/

	// 创建一个鼠标
	mouse := &Mouse{Name: "Logitech"}
	// 创建一个键盘
	keyboard := &Keyboard{Name: "Logitech"}
	// 调用测试函数
	testInterface(mouse)
	testInterface(keyboard)
}

// USB usb接口
type USB interface {
	start() // 开始工作
	end()   // 结束工作
}

// Mouse 鼠标
type Mouse struct {
	Name string
}

func (m *Mouse) start() {
	println(m.Name, "鼠标开始工作")
}

func (m *Mouse) end() {
	println(m.Name, "鼠标结束工作")
}

// Keyboard 键盘
type Keyboard struct {
	Name string
}

func (k *Keyboard) start() {
	println(k.Name, "键盘开始工作")
}

func (k *Keyboard) end() {
	println(k.Name, "键盘结束工作")
}

func testInterface(usb USB) {
	usb.start()
	usb.end()
}
