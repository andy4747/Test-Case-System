package routes

import (
	"net/http"

	"github.com/angeldhakal/testcase-ms/handlers"
	"github.com/angeldhakal/testcase-ms/middlewares"
	"github.com/gorilla/mux"
)

func MainRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	userHandler := handlers.NewUserHandler()

	//path prefixing for user
	user := router.PathPrefix("/users").Subrouter()
	user.HandleFunc("/register", middlewares.LoggerMiddleware(userHandler.AddUser)).Methods(http.MethodPost)
	user.HandleFunc("/login", middlewares.LoggerMiddleware(userHandler.SignInUser)).Methods(http.MethodPost)

	//path prefixing

	return router
}
