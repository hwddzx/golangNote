package main

import "fmt"

func main() {
	// 1. 访问元素
	arr := [5]int{1, 2, 3, 4}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("Array at index %d is %d\n", i, arr[i])
	}

	for i, v := range arr {
		fmt.Printf("Array at index %d is %d\n", i, v)
	}

	// 2. 修改元素
	array := [5]int{10, 20, 30, 40, 50}
	array[2] = 35
	fmt.Println(array)

	// 声明包含5个元素的的指向整数
	// 用整型指针初始化索引为0和1的数组元素
	a3 := [5]*int{0: new(int), 1: new(int)}
	fmt.Println(a3)
	*a3[0] = 1
	*a3[1] = 2
	fmt.Println(a3)

	var a1 [5]string
	a2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	a1 = a2
	fmt.Printf("%p\n", &a1)
	fmt.Printf("%p\n", &a2)
}
