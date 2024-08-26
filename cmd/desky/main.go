package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/eterline/desky/internal/server"
)

func main() {
	go func() {
		err := server.Start()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down the server...")
}
