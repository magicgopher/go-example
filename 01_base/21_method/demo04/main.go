package main

import "fmt"

func main() {
	// 方法值类型接收者
	c1 := &Count{10}
	result1 := c1.Increment1()
	fmt.Println(result1)
	fmt.Println(c1.value)

	fmt.Println("===============")

	// 方法指针类型接收者
	c2 := Count{10}
	result2 := c2.Increment2()
	fmt.Println(result2)
	fmt.Println(c2.value)
}

type Count struct {
	value int
}

func (c *Count) Increment1() int {
	c.value++ // 修改原始对象
	return c.value
}

func Increment1(c *Count) int {
	c.value++
	return c.value
}

func (c Count) Increment2() int {
	c.value++ // 只会修改副本，原对象不变
	return c.value
}

func Increment2(c Count) int {
	c.value++
	return c.value
}
