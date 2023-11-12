package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome web")
	// PerformGetRequest()
	// PerformPostJSONRequest()
	PerformPostFormRequest()

}

func PerformGetRequest() {
	const myurl = "http://localhost:3000/get"

	resp, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.StatusCode)
	fmt.Println("Content Length:", resp.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(resp.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount: ", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(string(content))
}

func PerformPostJSONRequest() {
	const myurl = "http://localhost:3000/post"

	requestBody := strings.NewReader(`
		{
			"courseName": "Golang",
			"prince": 0
		}
	`)
	resp, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:3000/postform"

	data := url.Values{}
	data.Add("name", "shubham")
	data.Add("surame", "pandey")

	resp, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(content))

}
