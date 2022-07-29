package main

import "fmt"

func main() {
	// 匿名函数
	func(num int) int {
		sum := 0
		for i := 1; i <= num; i++ {
			sum += i
		}
		fmt.Println(sum)
		return sum
	}(100)

	// 闭包
	fmt.Printf("%d\n", Add()(3))
	fmt.Printf("%d\n", Add2(6)(3))

	var f = adder()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))

	// nextNumber为一个函数, 函数 i 为 0
	nextNumber := getSequence()

	fmt.Printf("%d ", nextNumber())
	fmt.Printf("%d ", nextNumber())
	fmt.Printf("%d ", nextNumber())

	nextNumber1 := getSequence()
	fmt.Printf("%d ", nextNumber1())
	fmt.Printf("%d ", nextNumber1())
}

func Add() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Add2(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func adder() func(int) int {
	var x int
	return func(d int) int {
		fmt.Println("x = ", x, "d = ", d)
		x += d
		return x
	}
}

func getSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}