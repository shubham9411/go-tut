package main

import (
	"fmt"
	"log"
	"mongo/router"
	"net/http"
)

func main() {
	fmt.Println("Mongo DB APi")
	fmt.Println("Server is getting started...")

	r := router.Router()
	log.Fatal(http.ListenAndServe(":3000", r))
	fmt.Println("Listening at port 3000")
}
