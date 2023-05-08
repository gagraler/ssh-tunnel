package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 定义一个全局变量
var (
	configMap map[string]string
)

// 初始化配置文件
func InitConfig() {
	configMap = make(map[string]string)
	file, err := os.Open("./config/ssh-tunnel.conf")
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid line:", line)
			continue
		}

		key, value := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
		configMap[key] = value
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read file:", err)
		return
	}
}

/**
 * GetConfig 函数用于获取配置文件中指定 key 的值，并以字符串形式返回
 * @param: - key 配置文件中的 key 名称
 * @return: 字符串类型的 key 对应的 value 值
 */
func GetConfig(key string) string {
	return configMap[key]
}
