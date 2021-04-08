package main

import (
	"github.com/angeldhakal/testcase-ms/handlers"
	"github.com/angeldhakal/testcase-ms/models"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	err := models.CreateTablesIfNotExists(models.Connect())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Listening at http://localhost:8080/")
	log.Infoln("Necessary tables are present.")

	r := handlers.MainRouter()
	err = http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatal(err.Error())
	}
}
