package server

import (
	"flag"
	"net/http"

	"github.com/eterline/desky/internal/config"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func newServer(sessonStore sessions.Store, cfg config.Settings) *server {
	templs := paths{
		login:     "login.html",
		dashboard: "dashboard.html",
		notFound:  "404.html",
		docker:    "docker.html",
		proxmox:   "proxmox.html",
	}

	s := &server{
		router:       mux.NewRouter(),
		sessionStore: sessonStore,
		templates:    templs,
		configs:      cfg,
		cookieKey:    config.RandStringBytes(32),
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
	s.router.HandleFunc("/logout", s.goLogOut)

	private := s.router.PathPrefix("/dashboard").Subrouter()
	private.Use(s.authUser)
	private.HandleFunc("/panel", s.goDasboard)
	private.HandleFunc("/docker", s.goDocker)
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
