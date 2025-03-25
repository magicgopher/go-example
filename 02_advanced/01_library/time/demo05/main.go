package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("程序开始")

	// 设置 3 秒的超时时间
	timeout := 3 * time.Second
	timeoutChan := time.After(timeout)

	// 模拟一个耗时操作 (例如，从网络读取数据)
	operationComplete := make(chan bool) // 创建一个 channel 用于通知操作完成
	go func() {
		// 模拟耗时操作，这里暂停 5 秒
		time.Sleep(5 * time.Second)
		fmt.Println("耗时操作完成!")
		operationComplete <- true // 发送信号到 channel
	}()

	// 使用 select 监听超时和操作完成
	select {
	case <-operationComplete:
		fmt.Println("操作成功完成，没有超时!")
	case <-timeoutChan:
		fmt.Println("操作超时!")
	}

	fmt.Println("程序结束.")
}
