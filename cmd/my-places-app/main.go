package main

import (
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/mrDuderino/my-places-app/internal/app/handler"
	"github.com/mrDuderino/my-places-app/internal/app/repository"
	"github.com/mrDuderino/my-places-app/internal/app/service"
	"github.com/mrDuderino/my-places-app/internal/pkg/server"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatal(err)
	}

	if err := godotenv.Load("configs/.env"); err != nil {
		log.Fatal(err)
	}

	db, err := sqlx.Open("pgx", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		viper.Get("db.host"), viper.Get("db.port"), viper.Get("db.dbname"), viper.Get("db.username"),
		os.Getenv("DB_PASSWORD"), viper.Get("db.sslmode")))
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(viper.Get("port").(string), handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
