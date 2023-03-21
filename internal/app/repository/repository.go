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

type Dish interface {
	CreateDish(placeId int, dish models.Dish) (int, error)
	GetAllDishes(userId, placeId int) ([]models.Dish, error)
	GetById(userId, dishId int) (models.Dish, error)
	GetByName(userId int, dishName string) (models.Dish, error)
	Delete(userId, dishId int) error
	Update(userId int, dishId int, input models.UpdateDishInput) error
}

type DiscountCard interface {
	CreateDiscountCard(placeId int, card models.DiscountCard) (int, error)
	GetAllDiscountCards(userId, placeId int) ([]models.DiscountCard, error)
	GetById(userId, discountId int) (models.DiscountCard, error)
	Delete(userId, discountId int) error
}

type Repository struct {
	Authorization
	Place
	Dish
	DiscountCard
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		Place:         NewPlacesRepository(db),
		Dish:          NewDishRepository(db),
		DiscountCard:  NewDiscountCardRepository(db),
	}
}
