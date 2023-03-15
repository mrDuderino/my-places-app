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

func (s *PlaceService) GetAllPlaces(userId int) ([]models.Place, error) {
	return s.repos.GetAllPlaces(userId)
}

func (s *PlaceService) GetById(userId, placeId int) (models.Place, error) {
	return s.repos.GetById(userId, placeId)
}

func (s *PlaceService) GetByName(userId int, placeName string) (models.Place, error) {
	return s.repos.GetByName(userId, placeName)
}

func (s *PlaceService) Delete(userId, placeId int) error {
	return s.repos.Delete(userId, placeId)
}

func (s *PlaceService) Update(userId int, placeId int, input models.UpdatePlaceInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repos.Update(userId, placeId, input)
}
