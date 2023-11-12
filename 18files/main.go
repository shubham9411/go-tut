package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome files")
	content := "This is going to files"

	file, err := os.Create("./log.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println("The length is: ", length)
	defer file.Close()

	readFile("./log.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)

	fmt.Println("Date: ", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
