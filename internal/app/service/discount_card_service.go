package service

import (
	"github.com/mrDuderino/my-places-app/internal/app/models"
	"github.com/mrDuderino/my-places-app/internal/app/repository"
	"github.com/sirupsen/logrus"
)

type DiscountCardService struct {
	repos      repository.DiscountCard
	placeRepos repository.Place
}

func NewDiscountCardService(repos repository.DiscountCard, placeRepos repository.Place) *DiscountCardService {
	return &DiscountCardService{
		repos:      repos,
		placeRepos: placeRepos,
	}
}

func (s *DiscountCardService) CreateDiscountCard(userId int, placeId int, card models.DiscountCard) (int, error) {
	if _, err := s.placeRepos.GetById(userId, placeId); err != nil {
		logrus.Debugf("place with id=%d does not exist or does not belong to user: %s", placeId, err.Error())
		return 0, err
	}

	return s.repos.CreateDiscountCard(placeId, card)
}

func (s *DiscountCardService) GetAllDiscountCards(userId, placeId int) ([]models.DiscountCard, error) {
	return s.repos.GetAllDiscountCards(userId, placeId)
}

func (s *DiscountCardService) GetById(userId, discountId int) (models.DiscountCard, error) {
	return s.repos.GetById(userId, discountId)
}

func (s *DiscountCardService) Delete(userId, discountId int) error {
	return s.repos.Delete(userId, discountId)
}
