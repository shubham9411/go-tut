package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time study")

	presentTime := time.Now()

	fmt.Println(presentTime.Format("02-01-2006 Monday 15:04:05"))

	createdDate := time.Date(2020, time.August, 20, 10, 23, 0, 0, time.UTC)

	fmt.Println(createdDate.Format("02-01-2006 Monday"))
}
