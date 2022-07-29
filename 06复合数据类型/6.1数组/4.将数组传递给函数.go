package main

import "fmt"

func main() {
	// 数组是一个值类型, 所有的值类型变量在赋值和作为参数传递时都将产生一次复制动作.
	// 如果直接将数组作为函数的参数, 则在函数调用时数组会复制一份作为函数参数.
	// 因此, 在函数体中无法修改传入的数组内容
	var array [1e6]int
	foo1(array)

	foo2(&array)
}

func foo1(array [1e6]int) {
	fmt.Println(array)
}

func foo2(array *[1e6]int) {
	fmt.Println(*array)
}
