package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/eterline/desky/internal/logging"
	"github.com/eterline/desky/internal/server"
)

func main() {
	output := logging.InitLogOutput("logs/desky.log", true)
	log.SetOutput(output)
	go func() {
		err := server.Run()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down the server...")
}
