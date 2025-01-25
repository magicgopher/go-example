// package main 定义包名。
// Go 程序以包的形式组织代码，`main` 是一个特殊的包，表示这是一个可独立运行的程序。
package main

// import "fmt" 导入了 Go 的标准库 `fmt` 包
// `fmt` 提供格式化输入输出功能，比如打印字符串到控制台。
import "fmt"

// func main() 是程序的入口函数。
// Go 程序从 `main` 函数开始执行，`main` 函数是一个特殊的函数，没有参数也没有返回值。
func main() {
	// fmt.Println("hello world.") 调用 `fmt` 包中的 `Println` 函数
	// 在控制台打印字符串 "hello world." 并自动换行。
	fmt.Println("hello world.")
}
