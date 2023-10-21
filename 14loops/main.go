package main

import "fmt"

func main() {
	fmt.Println("loops in golang")

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 6 {
			goto lco
		}

		rougueValue++
	}

	// fmt.Println("rougue value is: ", rougueValue)

lco:
	fmt.Println("jumping at lco ", rougueValue)
}
