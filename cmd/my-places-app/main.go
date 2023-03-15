package main

import (
	"github.com/mrDuderino/my-places-app/internal/app/handler"
	"github.com/mrDuderino/my-places-app/internal/pkg/server"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(server.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
