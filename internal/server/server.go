package server

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"

	"github.com/eterline/desky/internal/config"
	"github.com/eterline/desky/pkg/ve"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func NewServer(sessonStore sessions.Store, cfg config.Settings, auth ve.Auth) *server {
	templs := paths{
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
		templates:     templs,
		configs:       cfg,
		cookieKey:     config.RandStringBytes(32),
		proxmoxClient: ve.Authenticate(auth),
	}

	s.configRouter()
	s.configApiRouter()
	return s
}

func (s *server) configRouter() {
	s.router.Use(loggingMiddleware)
	s.router.NotFoundHandler = http.HandlerFunc(s.goHome)

	s.router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	s.router.PathPrefix("/node/").Handler(http.StripPrefix("/node/", http.FileServer(http.Dir("./node_modules"))))
	s.router.HandleFunc("/login", s.goLogin)
	s.router.HandleFunc("/logout", s.goLogOut)

	content := s.router.PathPrefix("/static/").Subrouter()
	content.Use(s.authUser)
	content.PathPrefix("").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	private := s.router.PathPrefix("/dashboard").Subrouter()
	private.Use(s.authUser)
	private.HandleFunc("/ws", wsConnection)
	private.HandleFunc("/panel", s.goDasboard)
	private.HandleFunc("/tty", s.goTty)
	private.HandleFunc("/docker", s.goDocker)
	private.HandleFunc("/proxmox", s.goProxmox)
	private.HandleFunc("/monitor", s.goSysInfo)
	private.HandleFunc("/tty", s.goTty)
}

func (s *server) configApiRouter() {
	api := s.router.PathPrefix("/api").Subrouter()
	api.Use(s.authUser)
	api.HandleFunc("/system", s.apiSystem)
	api.HandleFunc("/pct/{id}/{cmd}", s.apiPct)
	api.HandleFunc("/qm/{id}/{cmd}", s.apiQm)
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
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
