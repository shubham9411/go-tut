package main

import "fmt"

const loginToken string = "jsdhhfjkmblcjkbds"

func main() {
	var username string = "Shubham"
	fmt.Println(username)
	fmt.Printf("The variable is of type %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("The variable is of type %T \n", isLoggedIn)

	var small int8 = 127
	fmt.Println(small)
	fmt.Printf("The variable is of type %T \n", small)

	var sample string
	fmt.Println(sample)
	fmt.Printf("The variable is of type %T \n", sample)

	fmt.Println(loginToken)

	users := 300
	fmt.Println(users)
}
