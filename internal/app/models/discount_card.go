package models

import (
	"errors"
	"time"
)

type DiscountCard struct {
	Id          int       `json:"id" db:"id"`
	Number      string    `json:"number" db:"number" binding:"required"`
	Description string    `json:"description" db:"description"`
	ValidFrom   time.Time `json:"valid_from" db:"valid_from"`
	ValidTo     time.Time `json:"valid_to" db:"valid_to"`
}

type PlaceDiscountCard struct {
	Id             int
	PlaceId        int `db:"place_id"`
	DiscountCardId int `db:"discount_card_id"`
}

type UpdateDiscountCardInput struct {
	Number      *string    `json:"number"`
	Description *string    `json:"description"`
	ValidFrom   *time.Time `json:"valid_from"`
	ValidTo     *time.Time `json:"valid_to"`
}

func (i *UpdateDiscountCardInput) Validate() error {
	if i.Number == nil && i.Description == nil && i.ValidTo == nil && i.ValidFrom == nil {
		return errors.New("empty update values")
	}
	return nil
}
