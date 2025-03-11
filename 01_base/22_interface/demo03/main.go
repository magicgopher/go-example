package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		接口类型断言是一种机制，用于从接口类型变量中提取其底层具体类型的值，或者检查接口变量是否持有某个特定的具体类型。

		格式：
			方式一：
				1. x := i.(T) //不安全会panic
				2. x, ok := i.(T) //安全
			方式二：
				switch i.(type) {
				case T:
					// 处理 T 类型
					...
				default:
					// 处理其他类型
				}
	*/

	// 创建不同类型的图形
	c := Circle{Radius: 5}
	sq := Square{Side: 4}
	t := Triangle{Side: 3}
	r := Rectangle{Width: 6, Height: 4}

	// 使用类型断言
	fmt.Println("使用类型断言")
	describeShape(c)
	describeShape(sq)
	describeShape(t)
	describeShape(r)

	fmt.Println("\n使用 switch 语句类型断言:")
	// 使用类型开关
	describeShapeSwitch(c)
	describeShapeSwitch(sq)
	describeShapeSwitch(t)
	describeShapeSwitch(r)
}

// Shape 图形接口
type Shape interface {
	Perimeter() float64 // 周长
	Area() float64      // 面积
}

// Circle 圆
type Circle struct {
	Radius float64 // 半径
}

// Perimeter 计算圆的周长
func (c Circle) Perimeter() float64 {
	return c.Radius * 2 * math.Pi
}

// Area 计算圆的面积
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Rectangle 矩形
type Rectangle struct {
	Width  float64 // 宽度
	Height float64 // 高度
}

// Perimeter 计算矩形的周长
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area 计算矩形的面积
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Square 正方形
type Square struct {
	Side float64
}

// Perimeter 计算正方形的周长
func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

// Area 计算正方形的面积
func (s Square) Area() float64 {
	return s.Side * s.Side
}

// Triangle 三角形
type Triangle struct {
	Side float64
}

// Perimeter 计算三角形的周长
func (t Triangle) Perimeter() float64 {
	return 3 * t.Side
}

// Area 计算三角形的面积
func (t Triangle) Area() float64 {
	return (math.Sqrt(3) / 4) * t.Side * t.Side
}

// 使用类型断言判断具体类型
func describeShape(s Shape) {
	// 使用类型断言
	if c, ok := s.(Circle); ok {
		fmt.Printf("这是一个圆形，半径为: %.2f, 周长: %.2f, 面积: %.2f\n", c.Radius, c.Perimeter(), c.Area())
		return
	}
	if sq, ok := s.(Square); ok {
		fmt.Printf("这是一个正方形，边长: %.2f, 周长: %.2f, 面积: %.2f\n", sq.Side, sq.Perimeter(), sq.Area())
		return
	}
	if t, ok := s.(Triangle); ok {
		fmt.Printf("这是一个三角形，边长: %.2f, 周长: %.2f, 面积: %.2f\n", t.Side, t.Perimeter(), t.Area())
		return
	}
	if r, ok := s.(Rectangle); ok {
		fmt.Printf("这是一个长方形，宽: %.2f, 高: %.2f, 周长: %.2f, 面积: %.2f\n", r.Width, r.Height, r.Perimeter(), r.Area())
		return
	}
	fmt.Println("未知形状类型")
}

// 使用类型开关判断具体类型
func describeShapeSwitch(s Shape) {
	switch v := s.(type) {
	case Circle:
		fmt.Printf("这是一个圆形，半径为: %.2f, 周长: %.2f, 面积: %.2f\n", v.Radius, v.Perimeter(), v.Area())
	case Square:
		fmt.Printf("这是一个正方形，边长: %.2f, 周长: %.2f, 面积: %.2f\n", v.Side, v.Perimeter(), v.Area())
	case Triangle:
		fmt.Printf("这是一个三角形，边长: %.2f, 周长: %.2f, 面积: %.2f\n", v.Side, v.Perimeter(), v.Area())
	case Rectangle:
		fmt.Printf("这是一个长方形，宽: %.2f, 高: %.2f, 周长: %.2f, 面积: %.2f\n", v.Width, v.Height, v.Perimeter(), v.Area())
	default:
		fmt.Println("未知形状类型")
	}
}
