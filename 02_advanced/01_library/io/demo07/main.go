package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	// 创建一个 bytes.Buffer，它实现了 io.Reader
	r := bytes.NewBuffer([]byte("Hello, World!"))

	// 使用 io.NopCloser 将其转换为 io.ReadCloser
	rc := io.NopCloser(r)

	// 现在，我们可以将 rc 传递给一个需要 io.ReadCloser 的函数
	buf := make([]byte, 10)
	n, err := rc.Read(buf)
	if err != nil && err != io.EOF {
		fmt.Println("读取错误:", err)
		return
	}
	fmt.Printf("读取了 %d 个字节: %s\n", n, buf[:n])

	// 关闭 rc，但它实际上什么也不做
	err = rc.Close()
	if err != nil {
		fmt.Println("关闭错误:", err)
	}
}
