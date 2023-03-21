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
	GetAllPlaces(userId int) ([]models.Place, error)
	GetById(userId, placeId int) (models.Place, error)
	GetByName(userId int, placeName string) (models.Place, error)
	Delete(userId, placeId int) error
	Update(userId int, placeId int, input models.UpdatePlaceInput) error
}

type Dish interface {
	CreateDish(userId int, placeId int, dish models.Dish) (int, error)
	GetAllDishes(userId, placeId int) ([]models.Dish, error)
	GetById(userId, placeId int) (models.Dish, error)
	GetByName(userId int, dishName string) (models.Dish, error)
	Delete(userId, dishId int) error
	Update(userId int, dishId int, input models.UpdateDishInput) error
}

type DiscountCard interface {
	CreateDiscountCard(userId int, placeId int, card models.DiscountCard) (int, error)
	GetAllDiscountCards(userId, placeId int) ([]models.DiscountCard, error)
	GetById(userId, discountId int) (models.DiscountCard, error)
	Delete(userId, discountId int) error
}

type Service struct {
	Authorization
	Place
	Dish
	DiscountCard
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Place:         NewPlaceService(repos.Place),
		Dish:          NewDishService(repos.Dish, repos.Place),
		DiscountCard:  NewDiscountCardService(repos.DiscountCard, repos.Place),
	}
}
