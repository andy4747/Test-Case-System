package middlewares

import (
    log "github.com/sirupsen/logrus"
    "net/http"
)

func LoggerMiddleware(f http.HandlerFunc) http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request) {
        log.Println(r.Host,"logger middleware")
        f(w,r)
    }
}
