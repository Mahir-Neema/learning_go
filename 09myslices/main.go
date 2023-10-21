package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("type pf fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	highScores = append(highScores, 855, 244, 444)

	sort.Ints(highScores)
	fmt.Println(highScores)

	// removing a vlue from slices based on index

	var courses = []string{"react", "js", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
