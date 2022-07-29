package main

import "fmt"

func main() {

	// 表达式switch
	grade := "B"
	marks := 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	fmt.Printf("你的成绩为%s\n", grade)

	// 类型switch
	var x interface{}
	x = 1
	switch i := x.(type) {
	case nil:
		fmt.Println(`nil`)
	case int:
		fmt.Println(`int`)
	case float64:
		fmt.Println(`float64`)
	default:
		fmt.Println(`other`)
		fmt.Println(i)
	}
}
