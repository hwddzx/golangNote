package main

import "fmt"

func main() {
	a := 11
	if a > 20 {
		fmt.Println("a大于20")
	} else if a < 10 {
		fmt.Println("a小于10")
	} else {
		fmt.Println("a大于10, 小于20")
	}
}
