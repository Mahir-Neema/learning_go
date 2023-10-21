package main

import "fmt"

func main() {
	fmt.Println("structs in golang")
	// no inheritance in golang no parent no child

	mahir := User{"mahir", "mahir@gmail.com", true, 20}
	fmt.Println(mahir)
	fmt.Printf("mahir details: %+v\n", mahir)
}

// Name, Email first captical letter ->> Public

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
