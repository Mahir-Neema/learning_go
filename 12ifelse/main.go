package main

import "fmt"

func main() {
	fmt.Println("if else in golang")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "regular user"
	} else if loginCount > 10 {
		result = "watch out"
	} else {
		result = "exact"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("number less than 10")
	} else {
		fmt.Println("number more than ten")
	}
}
