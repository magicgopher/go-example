package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	// Copy函数用于将数据从一个 io.Reader 复制到 io.Writer，直到源数据耗尽或者发生错误，它返回复制的字节数和可能遇到的错误。
	// 创建一个字符串作为输入源
	src := strings.NewReader("Hello, World!")
	// 创建一个缓冲区作为输出目标
	dst := new(bytes.Buffer)
	// 使用 io.Copy 函数将数据从 src 复制到 dst
	n, err := io.Copy(dst, src)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Copied %d bytes:%s\n", n, dst.String())
}
