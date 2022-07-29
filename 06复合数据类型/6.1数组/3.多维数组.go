package main

import "fmt"

func main() {
	// 数组本身只有一个维度, 但是可以组合多个数组创建多维数组. 多维数组多用于管理具有依赖关系的数据(例如坐标系等)
	var a1 [4][2]int
	fmt.Println(a1)

	a2 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	fmt.Println(a2)
}
