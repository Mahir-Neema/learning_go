package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("rating dedo ji")

	// comma ok syntax /  comma error syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks", input)
	fmt.Printf("type %T", input)
}
