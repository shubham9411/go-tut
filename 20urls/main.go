package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?courseName=react&payment=sdjhfbshjdfb"

func main() {
	fmt.Println("URLS")
	fmt.Println(myurl)

	// parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()

	fmt.Printf("Data type %T\n", qparams)

	fmt.Println(qparams["courseName"])

	for _, val := range qparams {
		fmt.Println(val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/learn",
		RawPath: "user=shubham",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
