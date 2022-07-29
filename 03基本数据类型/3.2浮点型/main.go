package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 浮点数
	var value1 float64
	value1 = 1
	value2 := 2
	value3 := 3.0
	v := value1 + value3
	fmt.Println(value1, value2, value3, v)
	fmt.Println(reflect.TypeOf(value1))
	fmt.Println(reflect.TypeOf(value2))
	fmt.Println(reflect.TypeOf(value3))
	fmt.Println(reflect.TypeOf(v))

	// 浮点数精度问题, 导致两个数相等
	value1 = 1
	value3 = 1.0000000000000001
	if value1 == value3 {
		fmt.Println("相等")
	}

	// 复数
	var v1 complex64 // 由两个float32构成的负数类型
	v1 = 3.2 + 12i
	v2 := 3.2 + 12i // 隐式声明, 默认是complex128类型
	v3 := complex(3.2, 12)
	v4 := v2 + v3 // 复数运算
	fmt.Println(reflect.TypeOf(v1))
	fmt.Println(reflect.TypeOf(v2))
	fmt.Println(reflect.TypeOf(v3))
	fmt.Println(v1, v2, v3, v4)
}
