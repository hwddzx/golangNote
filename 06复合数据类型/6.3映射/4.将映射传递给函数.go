package main

import "fmt"

func main() {
	// 创建一个映射, 存储颜色以及颜色对应的十六进制代码
	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}

	// 显示映射里的所有颜色
	for k, v := range colors {
		fmt.Printf("Key: %s Value: %s\n", k, v)
	}

	// 调用函数来移除指定的键
	removeColor(colors, "Coral")

	// 显示映射里的所有颜色
	for k, v := range colors {
		fmt.Printf("Key: %s Value: %s\n", k, v)
	}

}

func removeColor(colors map[string]string, color string) {
	delete(colors, color)
}
