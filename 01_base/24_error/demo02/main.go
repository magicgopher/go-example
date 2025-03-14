package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	/*
		简单示例
	*/

	file, err := os.Open("test.txt")
	if err != nil {
		//log.Fatal(err)
		//if PathErr, ok := err.(*os.PathError); ok {
		//	fmt.Println(PathErr.Op)
		//	fmt.Println(PathErr.Path)
		//	fmt.Println(PathErr.Err)
		//	return
		//}

		// Go1.13推出了errors包来对错误进行处理
		var pathErr *os.PathError
		if errors.As(err, &pathErr) {
			fmt.Println("Operation:", pathErr.Op)
			fmt.Println("Path:", pathErr.Path)
			fmt.Println("Error:", pathErr.Err)
			return
		}
		return
	}
	defer file.Close() // 关闭文件
	fmt.Println(file.Name(), "文件打开成功.")
}
