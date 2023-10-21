package main

import "fmt"

func main() {
	fmt.Println("functions in golang")

	result := adder(3, 5)
	fmt.Println("result: ", result)

	proresult := proAdder(3, 4, 5, 6, 6, 4, 3)
	fmt.Println("pro result: ", proresult)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}
