package main

import "fmt"

func main() {
	fmt.Println("array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "asdfpple"
	fruitList[3] = "Appfsale"

	fmt.Println("fruit list is:", fruitList)
	fmt.Println("lenght is:", len(fruitList))

	var vegList = [3]string{"potato", "cabbage", "spinach"}
	fmt.Println("vegs are:", vegList)
}
