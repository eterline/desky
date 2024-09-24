package server

import (
	"net/http"

	"github.com/eterline/desky/internal/config"
	"github.com/eterline/desky/pkg/ve"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type Routes map[string]http.HandlerFunc

func NewServer(sessonStore sessions.Store, cfg config.Settings, auth ve.Auth) *server {
	templates := paths{
		index:     "templates/index.html",
		login:     "templates/login.html",
		dashboard: "templates/dashboard.html",
		docker:    "templates/docker.html",
		proxmox:   "templates/proxmox.html",
		monitor:   "templates/monitor.html",
		tty:       "templates/tty.html",
	}

	s := &server{
		router:        mux.NewRouter(),
		sessionStore:  sessonStore,
		templates:     templates,
		configs:       cfg,
		cookieKey:     config.RandStringBytes(32),
		proxmoxClient: ve.Authenticate(auth),
	}

	s.configPagesRouter()
	s.configApiRouter()
	s.configPublicRouter()
	return s
}

func (s *server) configPagesRouter() {
	s.router.Use(loggingMiddleware)
	s.router.NotFoundHandler = http.HandlerFunc(s.Home)

	content := s.router.PathPrefix("/static/").Subrouter()
	content.Use(s.authUser)
	content.PathPrefix("").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	s.BuildSubRoute(
		Routes{
			"/ws":      wsConnection,
			"/panel":   s.Dasboard,
			"/tty":     s.Tty,
			"/docker":  s.Docker,
			"/proxmox": s.Proxmox,
			"/monitor": s.SysInfo,
		},
		"/dashboard", s.authUser,
	)
}

func (s *server) configApiRouter() {
	s.BuildSubRoute(
		Routes{
			"/system":         s.apiSystem,
			"/pct/{id}/{cmd}": s.apiPct,
			"/qm/{id}/{cmd}":  s.apiQm,
		},
		"/api", s.authUser,
	)
}

func (s *server) configPublicRouter() {
	s.dirHandleFiles("/public/", "/public/", "./public")
	s.dirHandleFiles("/node/", "/node/", "./node_modules")
	s.router.HandleFunc("/login", s.Login)
	s.router.HandleFunc("/logout", s.Logout)
}
