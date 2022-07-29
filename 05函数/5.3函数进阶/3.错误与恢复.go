package main

import "fmt"

func main() {
	test()
}

func test() {
	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		func(){
			recover()
		}()
	}()

	defer fmt.Println(recover())

	defer recover()

	panic("发生错误")
}