package main

import "fmt"

func main() {
	defer fmt.Println("ONE")
	defer fmt.Println("TWO")
	defer fmt.Println("THREE")
	fmt.Println("hello ji")
	// lifo order
}
