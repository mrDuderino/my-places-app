package main

import (
	"context"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	"github.com/mrDuderino/my-places-app/internal/app/handler"
	"github.com/mrDuderino/my-places-app/internal/app/repository"
	"github.com/mrDuderino/my-places-app/internal/app/service"
	"github.com/mrDuderino/my-places-app/internal/pkg/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	if err := godotenv.Load("configs/.env"); err != nil {
		logrus.Fatalf("error initializing evironment values: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		DbName:   viper.GetString("db.dbname"),
		User:     viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("error initializing db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	go func() {
		err := srv.Run(viper.GetString("port"), handlers.InitRoutes())
		if err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("error on runnning http server: %s", err.Error())
		}
	}()
	logrus.Infoln("my-places-app http server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logrus.Infoln("my-places-app http server is shutting down")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error on db connection closing: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
