package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrDuderino/my-places-app/internal/app/models"
)

const (
	UsersTable      = "users"
	PlacesTable     = "places"
	UserPlacesTable = "user_places"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type Place interface {
}

type Repository struct {
	Authorization
	Place
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}
