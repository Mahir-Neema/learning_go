package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("rating dedo ji")

	// comma ok syntax /  comma error syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("fsfsad", numRating+1)
	}
	// fmt.Printf("type %T", input)
}
