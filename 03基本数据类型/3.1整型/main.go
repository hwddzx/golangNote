package main

import (
	"fmt"
	"reflect"
)

var value1 int32 // 全局声明

func main() {
	value2 := 64 // 函数内部声明
	var value3 int64 = 100

	fmt.Println(reflect.TypeOf(value1))
	fmt.Println(reflect.TypeOf(value2))
	fmt.Println(reflect.TypeOf(value3))
	fmt.Println(reflect.TypeOf(value3) == reflect.TypeOf(value2))
}
