package server

import (
	"fmt"
	"net/http"

	"github.com/eterline/desky/internal/config"
	"github.com/eterline/desky/pkg/ve"
	"github.com/gorilla/sessions"
)

func Start() error {
	cfg := config.ParseSettings()
	ssStore := sessions.NewCookieStore([]byte(cfg.SessionStoreKey))
	proxmox := ve.InitBase(
		cfg.Proxmox.User,
		cfg.Proxmox.Password,
		cfg.Proxmox.Host,
		cfg.Proxmox.Port,
	)
	srv := NewServer(ssStore, cfg, &proxmox)
	config.PrintLogo(
		cfg.Proxmox.Up,
		cfg.Address.Ip,
		cfg.Address.Port,
		cfg.Auth,
	)
	return http.ListenAndServeTLS(
		fmt.Sprintf("%s:%s", cfg.Address.Ip, cfg.Address.Port),
		cfg.Tls.Crt,
		cfg.Tls.Key,
		srv.router,
	)
}
