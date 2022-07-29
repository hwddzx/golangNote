package main

import "fmt"

/*
	defer用于延迟调用指定函数, defer关键字只能出现在函数内部

	1. 只有当defer语句全部执行, defer所在函数才算真正结束执行.
	2. 当函数中有defer语句时, 需要等待所有defer语句执行完毕, 才会执行return语句

	因为defer的延迟特点, 可以把defer语句用于回收资源,清理收尾等工作.
*/

var i = 0

func print(i int) {
	fmt.Println(i)
}

func main() {
	for ; i < 5; i++ {
		defer print(i)
	}
}
