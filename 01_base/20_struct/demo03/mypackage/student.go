package mypackage

// 在 Go 语言中，结构体的可见性（也叫访问控制或导出性）是由字段和类型的命名规则决定的。
// Go 使用一种简单而直观的机制来控制可见性：首字母大写表示公开（exported），首字母小写表示私有（unexported）。
// 这种规则不仅适用于结构体字段，还适用于变量、函数、类型等。本节将详细讲解结构体的可见性，包括字段、方法和嵌套结构体的表现。

// Student1 结构体 其他包可见
type Student1 struct {
	Name string
	Age  int
}

// student2 结构体 其他包不可见
type student2 struct {
	name string
	age  int
}
