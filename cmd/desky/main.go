package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/eterline/desky/internal/config"
	"github.com/eterline/desky/internal/server"
	"github.com/eterline/desky/pkg/logging"
)

func main() {
	conf := config.ParseSettings()

	logging.InitLogger("logs", "desky.log")
	log := logging.ReturnEntry()

	srv := server.InitServer(conf, log.Logger)

	go func() {
		err := srv.Run()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutting down server")
}
