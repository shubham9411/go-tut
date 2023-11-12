package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Web requests")
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type %T", resp)

	defer resp.Body.Close()

	databytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var content string = string(databytes)
	fmt.Println("content: ", content)
}
