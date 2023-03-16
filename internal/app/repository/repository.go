package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrDuderino/my-places-app/internal/app/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type Place interface {
	CreatePlace(userId int, place models.Place) (int, error)
	GetAllPlaces(userId int) ([]models.Place, error)
	GetById(userId, placeId int) (models.Place, error)
	GetByName(userId int, placeName string) (models.Place, error)
	Delete(userId, placeId int) error
	Update(userId int, placeId int, input models.UpdatePlaceInput) error
}

type Repository struct {
	Authorization
	Place
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		Place:         NewPlacesRepository(db),
	}
}
