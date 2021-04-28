package middlewares

import "net/http"

func IsAuthed(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TODO check if the request is authorized
		next(w, r)
	}
}
