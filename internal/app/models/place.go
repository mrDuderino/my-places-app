package models

import "errors"

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

type UpdatePlaceInput struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Address     *string `json:"address"`
	Rating      *int    `json:"rating"`
}

func (i *UpdatePlaceInput) Validate() error {
	if i.Name == nil && i.Description == nil && i.Address == nil && i.Rating == nil {
		return errors.New("empty update values")
	}
	return nil
}
