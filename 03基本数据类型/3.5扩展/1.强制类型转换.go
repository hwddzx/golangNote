package main

import "fmt"

func main() {
	sum := 11
	count := 3
	mean := float32(sum) / float32(count)
	mean2 := sum / count
	fmt.Println(mean)
	fmt.Println(mean2)
}
