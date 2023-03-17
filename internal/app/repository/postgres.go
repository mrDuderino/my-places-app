package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	UsersTable       = "users"
	PlacesTable      = "places"
	UserPlacesTable  = "user_places"
	DishesTable      = "dishes"
	PlaceDishesTable = "place_dishes"
)

type Config struct {
	Host     string
	Port     string
	DbName   string
	User     string
	Password string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.DbName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
