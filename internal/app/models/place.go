package models

type Place struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name" binding:"required"`
	Description string `json:"description" db:"description"`
	Address     string `json:"address" db:"address"`
	Rating      int    `json:"rating" db:"rating"`
}

type UserPlace struct {
	Id      int
	UserId  int `db:"user_id"`
	PlaceId int `db:"place_id"`
}
