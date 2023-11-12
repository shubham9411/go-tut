package main

import "fmt"

func main() {
	fmt.Println("If else")

	loginCount := 29
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else {
		result = "Watch"
	}

	fmt.Printf("result %v\n", result)

	if 9%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is high")
	}
}
