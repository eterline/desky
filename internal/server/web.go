package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eterline/desky/internal/config"
	"github.com/gorilla/sessions"
)

func Start() error {
	settings := config.ParseSettings()
	ssStore := sessions.NewCookieStore([]byte(settings.SessionStoreKey))
	srv := newServer(ssStore, settings)
	log.Printf("Desky succsessfully started on: %s:%s", settings.Address.Ip, settings.Address.Port)
	return http.ListenAndServeTLS(
		fmt.Sprintf("%s:%s", settings.Address.Ip, settings.Address.Port),
		settings.Tls.Crt,
		settings.Tls.Key,
		srv.router,
	)
}
