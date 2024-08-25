package server

import (
	"flag"
	"net/http"

	"github.com/eterline/desky/internal/config"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type server struct {
	router       *mux.Router
	sessionStore sessions.Store
	templates    paths
	configs      config.Settings
}

type paths struct {
	login     string
	dashboard string
	notFound  string
}

func newServer(sessonStore sessions.Store, cfg config.Settings) *server {
	templs := paths{
		login:     "login.html",
		dashboard: "dashboard.html",
		notFound:  "404.html",
	}

	s := &server{
		router:       mux.NewRouter(),
		sessionStore: sessonStore,
		templates:    templs,
		configs:      cfg,
	}

	s.configRouter()
	s.configApiRouter()
	return s
}

func (s *server) configRouter() {
	var dir string
	flag.StringVar(&dir, "static", "./static", "static content path")
	flag.Parse()

	s.router.Use(loggingMiddleware)
	notF := http.HandlerFunc(s.goNotFound)
	s.router.NotFoundHandler = notF

	s.router.HandleFunc("/", s.goHome)
	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	s.router.HandleFunc("/login", s.goLogin)

	private := s.router.PathPrefix("/dashboard").Subrouter()
	private.Use(s.authUser)
	private.HandleFunc("/panel", s.goDasboard)
	private.HandleFunc("/docker", s.goDasboard)
	private.HandleFunc("/vm", s.goDasboard)
}

func (s *server) configApiRouter() {
	private := s.router.PathPrefix("/api").Subrouter()
	private.Use(s.authUser)
	private.HandleFunc("/system", s.apiSystem)
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code)
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int) {
	w.WriteHeader(code)
}

func chekUser(s config.Settings, u string, p string) bool {
	if s.User.Username == u && s.User.Password == p {
		return true
	}
	return false
}
