package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	// var ptr *int
	// fmt.Println("Value of ptr : ", ptr)

	myNumber := 23

	var ptr = &myNumber
	fmt.Println("Value of ptr : ", ptr)
	fmt.Println("Value of ptr : ", *ptr)

}
