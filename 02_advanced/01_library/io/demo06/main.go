package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// fun1
	//fun1()

	// fun2
	//fun2()

	// fun3
	fun3()
}

// 写入到标准输出
func fun1() {
	_, err := io.WriteString(os.Stdout, "Hello, World!\n")
	if err != nil {
		fmt.Println("写入错误:", err)
	}
}

// 写入到文件
func fun2() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("创建文件错误:", err)
		return
	}
	defer file.Close()

	_, err = io.WriteString(file, "This is a test string.\n")
	if err != nil {
		fmt.Println("写入文件错误:", err)
	}
}

// 写入到内存buffer
func fun3() {
	var buffer bytes.Buffer
	_, err := io.WriteString(&buffer, "This string is writed to buffer.\n")
	if err != nil {
		fmt.Println("写入buffer错误:", err)
	}
	fmt.Print(buffer.String())
}
