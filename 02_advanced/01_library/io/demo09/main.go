package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	// 创建一个 io.Reader 切片，包含两个 Reader：
	// 1. strings.NewReader 从字符串 "Hello, " 创建一个 Reader
	// 2. bytes.NewBufferString 从字符串 "World!!!" 创建一个 BufferReader
	readers := []io.Reader{
		strings.NewReader("Hello, "),
		bytes.NewBufferString("World!!!"),
	}

	// 使用 io.MultiReader 将多个 Reader 组合成一个单一的 Reader
	reader := io.MultiReader(readers...)

	// 创建一个字节切片，用于存储从 Reader 中读取的数据
	data := make([]byte, 0, 128) // 初始长度为 0，容量为 128

	// 创建一个字节切片，用作临时缓冲区
	buf := make([]byte, 10) // 缓冲区大小为 10 字节

	// 循环读取 Reader 中的数据，直到遇到 io.EOF（文件结束）
	for n, err := reader.Read(buf); err != io.EOF; n, err = reader.Read(buf) {
		// 如果读取过程中发生错误，则抛出 panic
		if err != nil {
			panic(err)
		}
		// 将读取到的数据追加到 data 切片中
		data = append(data, buf[:n]...) // buf[:n] 表示实际读取到的数据切片
	}

	// 将 data 切片转换为字符串并打印到控制台
	fmt.Printf("%s\n", data) // 输出：Hello, World!!!
}
