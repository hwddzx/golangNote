package main

import "fmt"

func main() {
	// for 语句后面的三个子句分别为: [初始化子语句], 条件子语句, [后置子语句]
	a := 0
	b := 5
	for a < b {
		a++
		fmt.Println(a)
	}
}
