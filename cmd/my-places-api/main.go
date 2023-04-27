package main

import (
	"context"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/mrDuderino/my-places-app/internal/app/handler"
	"github.com/mrDuderino/my-places-app/internal/app/repository"
	"github.com/mrDuderino/my-places-app/internal/app/service"
	"github.com/mrDuderino/my-places-app/internal/pkg/server"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const webPort = "80"

func main() {

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DbName:   os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  os.Getenv("SSL_MODE"),
	})
	if err != nil {
		logrus.Fatalf("error initializing db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	go func() {
		err := srv.Run(webPort, handlers.InitRoutes())
		if err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("error on runnning http server: %s", err.Error())
		}
	}()
	logrus.Infoln("my-places-api http server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logrus.Infoln("my-places-api http server is shutting down")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error on db connection closing: %s", err.Error())
	}
}
