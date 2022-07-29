package main

import (
	"fmt"
	"strconv"
)

// 这个包主要用于字符串与其他类型的转换
func main() {
	// 十进制转换为字符串
	fmt.Println(strconv.Itoa(123))

	// 64浮点数转换为字符串
	fmt.Println(strconv.FormatFloat(3.14, 'f', 1, 64))

	// 字符串转换为int
	fmt.Println(strconv.Atoi("123"))

	// 字符串转换为float64
	fmt.Println(strconv.ParseFloat("3.14", 1))
}
