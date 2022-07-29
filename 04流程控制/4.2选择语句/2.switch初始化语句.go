package main

import "fmt"

func main() {
	switch marks := 90; marks {
	case 90:
		fmt.Println("90")
	case 80:
		fmt.Println("80")
	default:
		fmt.Println("00")
	}
}
