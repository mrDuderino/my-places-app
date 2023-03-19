package service

import (
	"github.com/mrDuderino/my-places-app/internal/app/models"
	"github.com/mrDuderino/my-places-app/internal/app/repository"
)

type DishService struct {
	repos     repository.Dish
	placeRepo repository.Place
}

func NewDishService(repos repository.Dish, placeRepo repository.Place) *DishService {
	return &DishService{
		repos:     repos,
		placeRepo: placeRepo,
	}
}

func (s *DishService) CreateDish(userId int, placeId int, dish models.Dish) (int, error) {
	_, err := s.placeRepo.GetById(userId, placeId)
	if err != nil {
		return 0, err
	}
	return s.repos.CreateDish(placeId, dish)
}

func (s *DishService) GetAllDishes(userId, placeId int) ([]models.Dish, error) {
	_, err := s.placeRepo.GetById(userId, placeId)
	if err != nil {
		return nil, err
	}
	return s.repos.GetAllDishes(userId, placeId)
}

func (s *DishService) GetById(userId, dishId int) (models.Dish, error) {
	return s.repos.GetById(userId, dishId)
}

func (s *DishService) GetByName(userId int, dishName string) (models.Dish, error) {
	return s.repos.GetByName(userId, dishName)
}

func (s *DishService) Delete(userId, dishId int) error {
	return s.repos.Delete(userId, dishId)
}
