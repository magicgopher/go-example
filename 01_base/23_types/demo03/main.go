package main

import "fmt"

func main() {
	/*
		通过定义结构体和接口，封装数据和行为，提供更高的抽象层次。
	*/
	pool := DBPool{config: PoolConfig{MaxConnections: 10}}
	conn := MockConn{}
	result, _ := conn.Query("SELECT * FROM users")
	fmt.Printf("Pool config: %d connections, Query result: %s\n", pool.config.MaxConnections, result)
	// 输出: Pool config: 10 connections, Query result: Result of SELECT * FROM users
}

// 结构体：PoolConfig 表示连接池配置
type PoolConfig struct {
	MaxConnections int
}

// 接口：Connection 表示数据库连接行为
type Connection interface {
	Query(string) (string, error)
}

// 结构体：DBPool 表示连接池
type DBPool struct {
	config PoolConfig
}

// 实现 Connection 接口的简单类型
type MockConn struct{}

func (mc MockConn) Query(q string) (string, error) {
	return "Result of " + q, nil
}
