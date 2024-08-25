package main

import (
	"log"

	"github.com/eterline/desky/internal/server"
)

func main() {
	err := server.Start()
	if err != nil {
		log.Fatal(err.Error())
	}
}
