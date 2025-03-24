package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

func main() {
	// ReadAtLeast：从 io.Reader 中读取至少指定字节数的数据到缓冲区中。如果读取的字节数少于指定值，会返回错误
	// io.ReadAtLeast 函数提供了一种灵活的方式来读取数据，允许您指定必须读取的最小字节数。
	r := bytes.NewBufferString("Hello, World!")
	buf := make([]byte, 5)
	// 读取至少3个字节
	n, err := io.ReadAtLeast(r, buf, 3)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("读取了 %d 个字节: %s\n", n, string(buf[:n]))
	// 读取剩下的所有字节，因为剩下的字节数小于申请的buf的长度，所以buf会被填满，但是如果，要求的min大于剩下的字节数，就会报错。
	buf2 := make([]byte, 10)
	n, err = io.ReadAtLeast(r, buf2, 8)
	if err != nil {
		log.Println("Error:", err)
	} else {
		fmt.Printf("读取了 %d 个字节: %s\n", n, string(buf2[:n]))
	}
}
