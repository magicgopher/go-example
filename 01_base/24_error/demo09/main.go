package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config 表示配置文件的数据结构
type Config struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

// readConfigFromFile 从文件中读取配置，可能会触发 panic
func readConfigFromFile(filePath string) Config {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("无法打开文件: %v", err))
	}
	defer file.Close()

	// 读取并解析 JSON
	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		panic(fmt.Sprintf("JSON 解析失败: %v", err))
	}

	return config
}

// SafeLoadConfig 封装读取逻辑，使用 recover 捕获 panic
func SafeLoadConfig(filePath string) (config Config, err error) {
	defer func() {
		if r := recover(); r != nil {
			// 捕获 panic，返回默认配置和错误
			config = Config{} // 返回空配置
			err = fmt.Errorf("加载配置失败: %v", r)
		}
	}()

	// 调用可能触发 panic 的函数
	config = readConfigFromFile(filePath)
	err = nil // 如果没有 panic，错误为 nil
	return
}

func main() {
	// 示例 1：加载存在的配置文件
	config, err := SafeLoadConfig("config.json")
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Printf("成功加载配置: Host=%s, Port=%d\n", config.Host, config.Port)
	}

	// 示例 2：加载不存在的文件
	config, err = SafeLoadConfig("nonexistent.json")
	if err != nil {
		fmt.Println("错误:", err) // 输出类似: 错误: 加载配置失败: 无法打开文件: open nonexistent.json: no such file or directory
	} else {
		fmt.Printf("成功加载配置: Host=%s, Port=%d\n", config.Host, config.Port)
	}

	// 示例 3：加载格式错误的配置文件
	config, err = SafeLoadConfig("invalid.json")
	if err != nil {
		fmt.Println("错误:", err) // 输出类似: 错误: 加载配置失败: JSON 解析失败: invalid character 'i' looking for beginning of value
	} else {
		fmt.Printf("成功加载配置: Host=%s, Port=%d\n", config.Host, config.Port)
	}

	// 程序继续运行
	fmt.Println("程序正常结束")
}

// 模拟创建测试文件
func init() {
	// 创建有效的 config.json
	validConfig := Config{Host: "localhost", Port: 8080}
	data, _ := json.Marshal(validConfig)
	os.WriteFile("config.json", data, 0644)

	// 创建格式错误的 invalid.json
	os.WriteFile("invalid.json", []byte("invalid json"), 0644)
}
