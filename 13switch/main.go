package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Println("Dice number", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value 1 open")
		fallthrough
	case 2:
		fmt.Println("Dice value 2")
		fallthrough
	case 3:
		fmt.Println("Dice value 3")
	case 4:
		fmt.Println("Dice value 4")
	case 5:
		fmt.Println("Dice value 5")
	case 6:
		fmt.Println("Dice value 6 and roll again")
	}
}
