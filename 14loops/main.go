package main

import "fmt"

func main() {
	fmt.Println("LOoops")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println("days")

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}
	fmt.Println("---------")
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("index is %v value: %v\n", index, day)
	}

	value := 1

	for value < 10 {

		if value == 2 {
			goto lco
		}

		if value == 5 {
			value++
			continue
		}

		if value == 7 {
			break
		}

		fmt.Println("value", value)
		value++
	}

lco:
	fmt.Println("Jumpppp!!!!")
}
