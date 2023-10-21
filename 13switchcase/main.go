package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch case in go")
	// t := time.Now()
	// use .New() instead of .Seed()
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("vakue of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("one one one")
	case 2:
		fmt.Println("two two two")
	case 3:
		fmt.Println("three three three")
	case 4:
		fmt.Println("four four four")
	case 5:
		fmt.Println("five five five")
	case 6:
		fmt.Println("play again")
	default:
		fmt.Println("default case")
	}
}
