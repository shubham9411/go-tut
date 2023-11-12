package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs")

	// no inheritence super parent

	shubham := User{"Shubham", "shubham@go.dev", true, 22}
	fmt.Println(shubham)
	fmt.Printf("Shubham details are: %+v\n", shubham)
	fmt.Printf("%v details are: %+v\n", shubham.Name, shubham)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
