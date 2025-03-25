package main

import (
	"io"
	"os"
)

func main() {
	// 创建一个文件用于写入
	file, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 创建一个 io.Writer 切片，包含标准输出和文件
	writers := []io.Writer{
		os.Stdout,
		file,
	}

	// 使用 io.MultiWriter 将多个 Writer 组合成一个单一的 Writer
	multiWriter := io.MultiWriter(writers...)

	// 向 multiWriter 写入数据
	data := []byte("Hello, MultiWriter!")
	multiWriter.Write(data)
}
