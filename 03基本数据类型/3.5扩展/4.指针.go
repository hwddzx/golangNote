package main

import "fmt"

func main() {
	a := 10
	ap := &a
	fmt.Printf("a的地址: %x\n", &a)
	fmt.Printf("ap的地址: %x\n", ap)
	fmt.Printf("*ap的值: %d\n", *ap)

	// nil指针
	var ptr *int
	fmt.Printf("ptr的值是:%x\n", ptr)

	// 指针的指针
	var b *int // 定义一个a指针, 指向nil
	bP := &a   // 空指针a的指针
	fmt.Printf("b-->nil: %x\n", b)
	fmt.Printf("bP-->a: %x\n", bP)

	// 指针数组
	c := []int{10, 100, 200}
	var cP [3]*int
	for i := 0; i < len(c); i++ {
		cP[i] = &c[i]
		fmt.Printf("c[%d]的地址: %x\n", i, cP[i])
	}
	fmt.Println(cP)
	for i := 0; i < len(c); i++ {
		fmt.Printf("c[%d]的值: %d\n", i, *cP[i])
	}

	// 传递给函数
	x := 100
	y := 200
	fmt.Printf("交换之前x的值为: %d\n", x)
	fmt.Printf("交换之前y的值为: %d\n", y)
	swap(&x, &y)
	fmt.Printf("交换之后x的值为: %d\n", x)
	fmt.Printf("交换之后y的值为: %d\n", y)
}

func swap(x *int, y *int) {
	*x, *y = *y, *x
}
