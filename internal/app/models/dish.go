package models

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
