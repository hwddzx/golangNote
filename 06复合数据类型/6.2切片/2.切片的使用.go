package main

import "fmt"

func main() {
	s1 := []int{10, 20, 30, 40, 50}
	fmt.Println(s1)
	s1[1] = 25
	fmt.Println(s1)
	newS1 := s1[1:3]
	fmt.Println(newS1)
	newS2 := s1[0:2:5]
	fmt.Println(newS2)
	fmt.Println(len(newS2))
	fmt.Println(cap(newS2))

	// 切片扩容
	slice := []int{10, 20, 30, 40, 50}
	newSlice := slice[1:3]
	newSlice = append(newSlice, 60, 70, 80)
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

}
