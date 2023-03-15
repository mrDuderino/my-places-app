package service

import (
	"github.com/mrDuderino/my-places-app/internal/app/models"
	"github.com/mrDuderino/my-places-app/internal/app/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Place interface {
	CreatePlace(userId int, place models.Place) (int, error)
}

type Service struct {
	Authorization
	Place
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		Place:         NewPlaceService(repos),
	}
}
