package main

import "fmt"

func main() {
	// 1. 元素赋值
	colors := map[string]string{}
	colors["Red"] = "#DA1337"

	// 2. 查找与遍历
	value, exists := colors["Blue"]
	if exists {
		fmt.Println(value)
	}

	value = colors["Red"]
	if value != "" {
		fmt.Println(value)
	}

	color2 := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "＃ff7F50",
		"DarkGray":    "＃a9a9a9",
		"ForestGreen": "＃228b22",
	}

	for k, v := range color2 {
		fmt.Println(k, v)
	}

	// 3. 元素删除
	delete(color2, "Coral")
	fmt.Println(color2)
}
