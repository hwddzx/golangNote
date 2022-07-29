package main

import "fmt"

func main() {
	// break
	for i := 0; i < 10; i++ {
		fmt.Printf("i的值为: %d\n", i)
		if i > 4 {
			break
		}
	}

	// 打标签
LOOP1:
	for {
		x := 1
		switch {
		case x > 0:
			fmt.Println("A")
			break LOOP1
		case x == 1:
			fmt.Println("B")
		default:
			fmt.Println("C")
		}
	}

	// continue

	// goto
	var i int
	for {
		println(i)
		i++
		if i > 2 {
			goto BREAK
		}
	}

BREAK:
	println("break")
}
