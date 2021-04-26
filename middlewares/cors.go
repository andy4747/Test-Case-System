package middlewares

import (
    "net/http"
    "fmt"
)

func CORSMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request){
        fmt.Println("FROM CORS MIDDLEWARE")
        w.Header().Add("Access-Control-Allow-Origin","*")
        w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization, token, Set-Cookie")
        next(w,r)
    }
}
