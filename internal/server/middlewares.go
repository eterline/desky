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

func (s *server) authUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if s.configs.Auth {
			session, err := s.sessionStore.Get(r, SessionName)
			if err != nil {
				s.error(w, r, http.StatusInternalServerError, err)
				return
			}
			id, ok := session.Values["uID"]
			if !ok || id != s.cookieKey {
				http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
