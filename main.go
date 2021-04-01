package main

import (
	"github.com/angeldhakal/testcase-ms/handlers"
	"log"
	"net/http"
)

func main() {
	log.Println("Listening at http://localhost:8080/")

	r := handlers.NewRouter()
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatal(err.Error())
	}
}
