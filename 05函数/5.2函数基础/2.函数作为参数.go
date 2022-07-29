package main

import "fmt"

func pipe(ff func() bool) bool {
	return ff()
}

type FormatFunc func(s string, x, y int) string

func format(ff FormatFunc, s string, x, y int) string {
	return ff(s, x, y)
}

func main() {
	s1 := pipe(func() bool {
		return false
	})

	fmt.Println(s1)

	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d, %d", 10, 20)
	fmt.Println(s2)

}
