package main

import "fmt"

/*
	在使用递归时, 开发者需要设置退出条件, 否则递归将陷入无限循环
*/

// 斐波拉契数列
func fibonacci(n uint64) (result uint64) {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}

}

func main() {
	res := fibonacci(10)
	fmt.Println(res)
}
