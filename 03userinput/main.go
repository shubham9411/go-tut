package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome Back"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a random number")

	input, _ := reader.ReadString('\n')
	fmt.Println("You entered number: ", input)
	fmt.Printf("Type of number: %T", input)
}
