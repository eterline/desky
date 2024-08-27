package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/eterline/desky/internal/server"
)

func main() {
	logPath := fmt.Sprintf(
		"logs/desky_%v_%v_%v-%v_%v_%v.log",
		time.Now().Year(), time.Now().Month(), time.Now().Day(),
		time.Now().Hour(), time.Now().Minute(), time.Now().Second(),
	)
	file, _ := os.Create(logPath)
	log.SetOutput(file)
	defer file.Close()

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
