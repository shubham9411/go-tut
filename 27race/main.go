package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race conditions")

	var score = []int{0}

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	wg.Add(4)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("One routine")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()

		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Two routine")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()

		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Three routine")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()

		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Three routines")
		mut.RLock()
		for _, v := range score {
			fmt.Println("Score", v)
		}
		mut.RUnlock()

		wg.Done()
	}(wg, mut)

	wg.Wait()

	fmt.Println(score)
}
