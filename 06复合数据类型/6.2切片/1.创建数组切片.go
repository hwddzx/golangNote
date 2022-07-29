package main

import (
	"fmt"
)

func main() {
	// 1. make和切片字面量

	// 创建一个字符串切片, 其长度和容量都是5个元素
	s1 := make([]string, 5)
	fmt.Println(s1)

	// 创建一个整型切片, 其长度为3, 容量为5
	s2 := make([]int, 3, 5)
	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))

	s3 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	s4 := []int{1, 2, 3}
	fmt.Println(s3, s4)

	// 使用空字符串初始化第100个元素
	s5 := []string{99: "1"}
	fmt.Println(s5)


	// 2. nil和空切片
	//slice := make([]int, 0)
	var slice []int
	fmt.Println(slice)
}
