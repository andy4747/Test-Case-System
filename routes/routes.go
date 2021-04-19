package routes

import (
	"github.com/angeldhakal/testcase-ms/handlers"
	"github.com/angeldhakal/testcase-ms/middlewares"
	"github.com/gorilla/mux"
)

func MainRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	userHandler := handlers.NewUserHandler()

	//path prefixing for user
	user := router.PathPrefix("/users").Subrouter()
	user.HandleFunc("/register", middlewares.CORSMiddleware(middlewares.LoggerMiddleware(userHandler.AddUser)))
	user.HandleFunc("/login", middlewares.CORSMiddleware(middlewares.LoggerMiddleware(userHandler.SignInUser)))

	//path prefixing

	return router
}
