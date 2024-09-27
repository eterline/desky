package server

import (
	"log"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if _, ok := recover().(error); ok {
				log.Printf("[%v] %s > %s.", r.Response.StatusCode, r.RemoteAddr, http.StatusInternalServerError)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
		log.Printf(" %s > %s.", r.RemoteAddr, r.RequestURI)
	})
}
