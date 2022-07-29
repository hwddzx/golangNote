package main

import "fmt"

/*

Go 语言 中让人颇为称道的 个设计就是延迟（ ）语句，开发者可以在函数中添
加多个 语句。当函数执行到最后时（return 语句执行之前），这些 efer 语句会按照
“逆序”执行，最后该函数才退出

*/

func main() {
	fmt.Println("return:", a())

	fmt.Println("return:", b())
}

func a() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()

	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return i
}

func b() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()

	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return i
}