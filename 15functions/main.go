package main

import "fmt"

func main() {
	fmt.Println("Welcome functions")
	greeter()
	fmt.Println(adder(1, 2))
	fmt.Println(proadder(1, 2, 2, 3, 3))
}

func greeter() {
	fmt.Println("Namaste golang")
}

func adder(a int, b int) int {
	return a + b
}

func proadder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "result"
}
