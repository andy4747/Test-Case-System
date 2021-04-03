package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler).Methods(http.MethodGet)
	router.HandleFunc("/login",loginUser).Methods(http.MethodPost)
	router.HandleFunc("/register", registerUser).Methods(http.MethodPost)
	router.HandleFunc("/case/create", createTestCase).Methods(http.MethodPost)
	router.HandleFunc("/cases",getTestCases).Methods(http.MethodGet)

	return router
}
