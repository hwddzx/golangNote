package main

import "fmt"

func main() {
	str := `abc`
	for i, char := range str {
		fmt.Printf("第%d个字符的值为%s\n", i, string(char))
	}

	// 忽略index
	for _, char := range str {
		fmt.Println(char)
	}

	// 忽略第二个值
	for i := range str {
		fmt.Println(i)
	}

	// 忽略全部
	for range str {
		fmt.Println("执行成功")
	}

	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		fmt.Println(k, v)
	}

	numbers := []int{1, 2, 3, 4}
	for i, x := range numbers {
		fmt.Println(i, x)
	}
}
