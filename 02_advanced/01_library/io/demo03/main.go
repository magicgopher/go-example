package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	// CopyBuffer：是 io.Copy 的变种，允许用户指定一个缓冲区来减少底层系统调用的开销，如果缓冲区为 nil，则会自动分配一个 32KB 大小的缓冲区。
	src := strings.NewReader("Hello, World!")
	dst := new(bytes.Buffer)
	buf := make([]byte, 4) // 使用 4 字节的缓冲区

	n, err := io.CopyBuffer(dst, src, buf)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Copied %d bytes: %s\n", n, dst.String())
}
