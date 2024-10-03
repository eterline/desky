package server

import (
	"net/http"
)

func (s *server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if _, ok := recover().(error); ok {
				log.Printf("[%v] %s > %s.", r.Response.StatusCode, r.RemoteAddr)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
		log.Printf("%s > %s.", r.RemoteAddr, r.RequestURI)
	})
}

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Add("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
