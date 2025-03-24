package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// 创建多个字符串 Reader
	r1 := strings.NewReader("Hello, ")
	r2 := strings.NewReader("World!")

	// 创建 MultiReader，连接多个 Reader
	mr := io.MultiReader(r1, r2)

	// 读取数据，确保缓冲区足够大
	buf := make([]byte, 20)
	n, err := mr.Read(buf)

	if err != nil && err != io.EOF {
		fmt.Println("读取错误:", err)
		return
	}

	// 打印读取的字节数和数据
	fmt.Printf("读取了 %d 个字节: %s\n", n, buf[:n])
}
