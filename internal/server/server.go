package server

import (
	"fmt"
	"net/http"

	"github.com/eterline/desky/internal/config"
	"github.com/eterline/desky/pkg/notification"
	"github.com/eterline/desky/pkg/ve"
	opnsenseapi "github.com/eterline/opnsense-api"
	"github.com/gorilla/mux"
	"github.com/luthermonson/go-proxmox"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func (s *server) Run() error {
	log.Infof("Server started at: %s:%s", s.configs.Address.Ip, s.configs.Address.Port)
	return ListenConnections(s.configs, s.router)
}

func InitServer(cfg config.Settings, logger *logrus.Logger) *server {
	log = logger

	templates := map[string]string{
		"index":     "templates/index.html",
		"dashboard": "templates/dashboard.html",
		"proxmox":   "templates/proxmox.html",
		"monitor":   "templates/monitor.html",
		"systemd":   "templates/systemd.html",
		"tty":       "templates/tty.html",
		"opnsense":  "templates/opnsense.html",
	}

	var accountsProxm []*proxmox.Client
	for _, j := range cfg.Proxmox.Nodes {
		auth := ve.InitBase(j.User, j.Password, j.Host, j.Port)
		accountsProxm = append(accountsProxm, ve.Authenticate(&auth))
	}

	opnClinet, err := opnsenseapi.NewClient(
		cfg.Opnsense.Key,
		cfg.Opnsense.Secret,
		cfg.Opnsense.Host,
	)
	if err != nil {
		log.Panic(err)
	}

	s := &server{
		router:         mux.NewRouter(),
		templates:      templates,
		configs:        cfg,
		proxmoxClient:  accountsProxm,
		opnsenseClient: opnClinet,
		logger:         logger,
	}

	s.router.Use(s.loggingMiddleware)
	s.configPrivateStatic()
	s.configPagesRouter()
	s.configApiRouter()
	s.configPublicRouter()
	Notify(s.configs)
	return s
}

func ListenConnections(config config.Settings, router *mux.Router) error {
	if config.Tls.Enable {
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
