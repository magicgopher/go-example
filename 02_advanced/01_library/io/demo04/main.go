package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

func main() {
	// ReadFull：从 io.Reader 中读取正好指定字节数的数据，如果无法读取足够的字节，会返回ErrUnexpectedEOF错误
	// 创建一个 bytes.Buffer 模拟一个 io.
	r := bytes.NewBuffer([]byte("Hello, World!"))
	// 创建一个字节切片，用于存储读取的数据（缓冲区）
	buf1 := make([]byte, 5)
	// 尝试读取完整的5个字节
	n, err := io.ReadFull(r, buf1)
	if err != nil {
		log.Fatal("Error:", err)
		return
	}
	fmt.Printf("读取了 %d 个字节: %s\n", n, buf1)
	// 尝试读取完整的10个字节，但是剩下8个字节，不满足完整的10字节长度会出现ErrUnexpectedEOF错误
	buf2 := make([]byte, 10)
	n, err = io.ReadFull(r, buf2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Printf("读取了 %d 个字节: %s\n", n, buf2)
	}
}
