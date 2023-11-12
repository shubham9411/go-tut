package main

import "fmt"

func main() {
	fmt.Println("Maps:")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println(languages)
	fmt.Println(languages["JS"])
	delete(languages, "RB")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
