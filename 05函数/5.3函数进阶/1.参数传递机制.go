package main

import "fmt"

/*
值传递:
	值传递是传递参数的副本, 函数接收参数副本后, 使用变量的过程中可能对副本的值进行更改, 但不会影响原本的变量

引用传递:
	引用传递是将参数的指针传递给函数, 函数修改这个指针的值, 就会修改到原本参数的值

	好处:
	1. 使得多个函数能操作同一个对象
	2. 传指针比较轻量
	3. 赋予了函数直接修改外部变量的能力, 所以被修改的变量不再需要return返回
*/

func main() {
	x := 3
	fmt.Println("x = ", x, "&x = ", &x)

	y := add(x)
	fmt.Println("x = ", x, "y = ", y)

	z := addP(&x)
	fmt.Println("x = ", x, "z = ", z)
	fmt.Println("&x = ", &x)

}

func add(a int) int {
	a++
	return a
}

func addP(a *int) int {
	*a++
	return *a
}
