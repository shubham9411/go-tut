package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs")

	// no inheritence super parent

	shubham := User{"Shubham", "shubham@go.dev", true, 22}
	fmt.Println(shubham)
	fmt.Printf("Shubham details are: %+v\n", shubham)
	fmt.Printf("%v details are: %+v\n", shubham.Name, shubham)

	shubham.GetStatus()
	shubham.NewMail()
	fmt.Printf("%v\n", shubham.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "yolo@go.dev"
	fmt.Println("Email of user: ", u.Email)
}
