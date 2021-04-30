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

	//case handler
	caseHandler := handlers.NewCaseHandler()

	//path prefixing for cases
	cases := router.PathPrefix("/api/v1/cases").Subrouter()
	cases.HandleFunc("/get/{id}", caseHandler.GetCase)
	// cases.HandleFunc("/create", caseHandler.CreateTestCase)

	return router
}
