package main

import (
	"net/http"

	"github.com/angeldhakal/testcase-ms/routes"
	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Println("Listening at http://localhost:8080/")
	r := routes.MainRouter()
    headers := handlers.AllowedHeaders([]string{"X-Requested-With","Content-Type", "Authorization", "token", "Set-Cookie"})
    methods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "PUT", "DELETE"})
    origins := handlers.AllowedOrigins([]string{"*"})
	err := http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(r))

	if err != nil {
		log.Fatal(err.Error())
	}
}
