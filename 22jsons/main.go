package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome jsons")
	// EncodeJson()
	DecodeJSON()
}

func EncodeJson() {
	lcocourse := []course{
		{"Reactjs", 199, "sp.in", "abcabc", []string{"web-dev", "js"}},
		{"Angular", 499, "sp.in", "123455", []string{"web-dev", "js"}},
		{"Go", 499, "sp.in", "123455", nil},
	}

	finalJson, err := json.MarshalIndent(lcocourse, "", "  ")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", string(finalJson))

}

func DecodeJSON() {
	jsonData := []byte(`
		{
			"coursename": "Reactjs",
			"price": 199,
			"website": "sp.in",
			"tags": ["web-dev", "js"]
		}
	`)
	var lcocourse course

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("JSON was valid!!")
		json.Unmarshal(jsonData, &lcocourse)
		fmt.Printf("%#v\n", lcocourse)
	} else {
		fmt.Println("JSON was not valid!!!")
	}

	var myonlinedata map[string]interface{}

	json.Unmarshal(jsonData, &myonlinedata)
	fmt.Printf("%#v\n", myonlinedata)

	for k, v := range myonlinedata {
		fmt.Printf("Key is %v value is %v and type is %T\n", k, v, v)
	}
}
