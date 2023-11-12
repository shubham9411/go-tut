package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex

var signals = []string{"test"}

func main() {
	fmt.Println("Welcome go routines")
	// go greeter("Hello")
	// greeter("World")

	websitelist := []string{
		"https://google.com",
		"https://shubhampandey.in",
		"https://github.com",
		"https://go.dev",
	}

	for _, v := range websitelist {
		go getStatusCode(v)
		wg.Add(1)
	}

	wg.Wait()

	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("error in endpoint")
	}

	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()

	fmt.Printf("%d Status code for %s\n", res.StatusCode, endpoint)
}
