package server

import (
	"fmt"
	"net/http"

	"github.com/eterline/desky/internal/config"
	"github.com/gorilla/mux"
)

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
