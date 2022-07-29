package main

import "fmt"

func main() {
	// 手动填写参数
	age := ageMinOrMax("min", 1, 3, 2, 0)
	fmt.Printf("年龄最小的%d\n", age)
	// 数组作为参数
	ageArr := []int{7, 9, 3, 5, 1}
	// 将数组解包
	age = ageMinOrMax("max", ageArr...)
	fmt.Printf("年龄最大的%d\n", age)

	f1(ageArr...)
}

func ageMinOrMax(m string, a ...int) int {
	if len(a) == 0 {
		return 0
	}
	if m == "max" {
		max := a[0]
		for _, v := range a {
			if v > max {
				max = v
			}
		}
		return max
	} else if m == "min" {
		min := a[0]
		for _, v := range a {
			if v < min {
				min = v
			}
		}
		return min
	} else {
		e := -1
		return e
	}
}

func f1(arr ...int) {
	f2(arr...)
	fmt.Println("")
	f3(arr)
}

func f2(arr ...int) {
	for _, char := range arr {
		fmt.Printf("%d ", char)
	}
}

func f3(arr []int) {
	for _, char := range arr {
		fmt.Printf("%d ", char)
	}
}
