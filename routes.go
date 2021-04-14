package handlers

import (
	"github.com/angeldhakal/testcase-ms/handlers/cases_handler"
	"github.com/angeldhakal/testcase-ms/handlers/users_handler"
	"github.com/gorilla/mux"
	"net/http"
)

func MainRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler).Methods(http.MethodGet)
	router.HandleFunc("/login", users_handler.LoginUser).Methods(http.MethodPost)
	router.HandleFunc("/register", users_handler.RegisterUser).Methods(http.MethodPost)
	router.HandleFunc("/case/create", cases_handler.CreateTestCase).Methods(http.MethodPost)
	router.HandleFunc("/cases", cases_handler.GetTestCases).Methods(http.MethodGet)
	router.HandleFunc("/case/delete/{id}", cases_handler.DeleteTestCase).Methods(http.MethodDelete)
	router.HandleFunc("/case/{id}", cases_handler.GetTestCase).Methods(http.MethodGet)

	return router
}
