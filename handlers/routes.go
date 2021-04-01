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
	router.HandleFunc("/case/create", createTestPostCase).Methods(http.MethodPost)

	return router
}
