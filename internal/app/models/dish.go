package models

import "errors"

type Dish struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name" binding:"required"`
	Description string `json:"description" db:"description"`
	Rating      int    `json:"rating" db:"rating"`
}

type PlaceDish struct {
	Id      int
	PlaceId int `db:"place_id"`
	DishId  int `db:"dish_id"`
}

type UpdateDishInput struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Rating      *int    `json:"rating"`
}

func (i *UpdateDishInput) Validate() error {
	if i.Name == nil && i.Description == nil && i.Rating == nil {
		return errors.New("empty update values")
	}
	return nil
}
