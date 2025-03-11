package main

import "fmt"

func main() {
	/*
		通过 type 关键字来自定义类型，避免直接使用基础类型带来的误用问题。

		场景：区分用户ID和订单ID
	*/

	var uid UserID = 1001
	var oid OrderID = 50001
	getOrder(uid, oid) // 输出: UserID: 1001, OrderID: 50001
	// getOrder(oid, uid) // 编译错误：类型不匹配
}

type UserID int  // 用户ID
type OrderID int // 订单ID

func getOrder(uid UserID, oid OrderID) {
	fmt.Printf("UserID: %d, OrderID: %d\n", uid, oid)
}
