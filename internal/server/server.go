package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eterline/desky/internal/config"
	"github.com/eterline/desky/pkg/notification"
	"github.com/eterline/desky/pkg/ve"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func Run() error {
	cfg := config.ParseSettings()
	ssStore := sessions.NewCookieStore([]byte(cfg.SessionStoreKey))
	proxmox := ve.InitBase(
		cfg.Proxmox.User,
		cfg.Proxmox.Password,
		cfg.Proxmox.Host,
		cfg.Proxmox.Port,
	)
	srv := NewServer(ssStore, cfg, &proxmox)

	cfg.PrintLogo()

	return ListenConnections(cfg.Tls.Enable, cfg, srv.router)
}

func NewServer(sessonStore sessions.Store, cfg config.Settings, auth ve.Auth) *server {
	templates := paths{
		index:     "templates/index.html",
		dashboard: "templates/dashboard.html",
		docker:    "templates/docker.html",
		proxmox:   "templates/proxmox.html",
		monitor:   "templates/monitor.html",
		tty:       "templates/tty.html",
	}

	s := &server{
		router:        mux.NewRouter(),
		templates:     templates,
		configs:       cfg,
		proxmoxClient: ve.Authenticate(auth),
	}

	s.router.Use(loggingMiddleware)
	s.configPrivateStatic()
	s.configPagesRouter()
	s.configApiRouter()
	s.configPublicRouter()
	Notify(s.configs)
	return s
}

func ListenConnections(tls bool, config config.Settings, router *mux.Router) error {
	if tls {
		return http.ListenAndServeTLS(
			fmt.Sprintf("%s:%s", config.Address.Ip, config.Address.Port),
			config.Tls.Crt,
			config.Tls.Key,
			router,
		)
	}
	return http.ListenAndServe(
		fmt.Sprintf("%s:%s", config.Address.Ip, config.Address.Port),
		router,
	)

}

func Notify(c config.Settings) {
	gotify, err := notification.InitGotify(
		c.Notifications.Gotify.URL,
		c.Notifications.Gotify.KEY,
	)
	if err != nil {
		log.Println("failed send notification")
		return
	}
	err = gotify.Send("Desky panel", "Desky panel has been started.", 2)
	if err != nil {
		log.Println("failed send notification")
		return
	}
	log.Println("startup notification sent")
}
