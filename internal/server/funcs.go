package server

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"

	"github.com/eterline/desky/internal/config"
	"github.com/gorilla/mux"
)

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	log.Println(err)
	s.respond(w, r, code)
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int) {
	w.WriteHeader(code)
}

func chekUser(s config.Settings, u string, p string) bool {
	h := sha256.New()
	h.Write([]byte(p))
	hex.EncodeToString(h.Sum(nil))
	return s.User.Username == u && s.User.Password == ToHash(p)
}

func ToHash(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func errLog(err error) {
	if err != nil {
		log.Println(TEMPLATE_ERR, err.Error())
	}
}

// Build subrouter with middlewares <-- map["route"]HandleFunc, "/main", [...]middleware
func (s *server) BuildSubRoute(routes Routes, path string, middlewares ...mux.MiddlewareFunc) *mux.Router {
	sub := s.router.PathPrefix(path).Subrouter()

	for route, handler := range routes {
		sub.HandleFunc(route, handler)
	}

	if len(middlewares) > 0 {
		sub.Use(middlewares...)
	}
	return sub
}

func (s *server) dirHandleFiles(path, route, dir string) *mux.Route {
	filesDir := http.StripPrefix(route, http.FileServer(http.Dir(dir)))
	return s.router.PathPrefix(path).Handler(filesDir)
}
