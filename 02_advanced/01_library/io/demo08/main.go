package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// 创建一个字符串 Reader
	s := strings.NewReader("Hello, Golang!")

	// 创建一个LimitReader限制读取5个字节
	lr := io.LimitReader(s, 5)

	// 读取数据
	buf := make([]byte, 10)
	n, err := lr.Read(buf)

	if err != nil && err != io.EOF {
		fmt.Println("读取错误:", err)
		return
	}

	// 打印读取的字节数和数据
	fmt.Printf("读取了 %d 个字节: %s\n", n, buf[:n])

	//尝试再次读取
	n, err = lr.Read(buf)
	if err == io.EOF {
		fmt.Println("读取完成")
	}
}
