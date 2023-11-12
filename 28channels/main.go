package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome channels")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// myCh <- 5

	wg.Add(2)

	// Recieve only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-ch

		fmt.Println(isChannelOpen)
		fmt.Println(val)
		fmt.Println(<-ch)

		wg.Done()
	}(myCh, wg)

	// Send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 5
		close(ch)
		// ch <- 6

		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
