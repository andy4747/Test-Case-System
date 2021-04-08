package users_handler


import (
	"fmt"
	"net/http"
)

func registerUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	fmt.Fprintf(w,`{"endpoint":"register"}`)
}