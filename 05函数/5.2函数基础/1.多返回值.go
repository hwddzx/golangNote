package main

import "fmt"

func SumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}

func main() {
	x := 3
	y := 4

	xPLUSy, xTIMEy := SumAndProduct(x, y)
	fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
	fmt.Printf("%d * %d = %d\n", x, y, xTIMEy)

}
