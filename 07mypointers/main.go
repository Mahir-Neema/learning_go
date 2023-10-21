package main

import "fmt"

func main() {
	fmt.Println("welcome")

	// var ptr *int
	myNumber := 23
	var ptr = &myNumber
	fmt.Println(myNumber)
	fmt.Println(ptr)
	fmt.Println(*ptr)
}
