package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruits [4]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[3] = "Orange"

	fmt.Println(fruits)

	var fruitList = []string{"Apple", "Banana", "Orange"}

	fmt.Println(fruitList)

	fruitList = append(fruitList, "Mango")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)
	highScore[0] = 123
	highScore[1] = 945
	highScore[2] = 343
	highScore[3] = 455
	// highScore[4] = 455

	highScore = append(highScore, 555, 443)

	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	var courses = []string{"react", "ruby", "go", "cpp"}
	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
