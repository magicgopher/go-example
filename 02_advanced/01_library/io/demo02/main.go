package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	// CopyN：与 Copy类似，但它只复制指定数量的字节（n）。如果源数据不足 n 个字节，会返回 io.EOF
	src := strings.NewReader("Hello, World!")
	dst := new(bytes.Buffer)
	// 只复制前 5 个字节（包括5）
	n, err := io.CopyN(dst, src, 5)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Copied %d bytes:%s\n", n, dst.String())
}
