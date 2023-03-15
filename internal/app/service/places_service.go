package service

import (
	"github.com/mrDuderino/my-places-app/internal/app/models"
	"github.com/mrDuderino/my-places-app/internal/app/repository"
)

type PlaceService struct {
	repos *repository.Repository
}

func NewPlaceService(repos *repository.Repository) *PlaceService {
	return &PlaceService{repos: repos}
}

func (s *PlaceService) CreatePlace(userId int, place models.Place) (int, error) {
	return s.repos.CreatePlace(userId, place)
}
