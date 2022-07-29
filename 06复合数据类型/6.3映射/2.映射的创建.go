package main

import "fmt"

func main() {
	// 创建一个映射, 键的类型是string, 值的类型是int
	dict1 := make(map[string]int)

	// 创建一个映射, 并初始化
	dict2 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}

	fmt.Println(dict1)
	fmt.Println(dict2)
}
