package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/mrDuderino/my-places-app/internal/app/models"
)

type PlacesRepository struct {
	db *sqlx.DB
}

func NewPlacesRepository(db *sqlx.DB) *PlacesRepository {
	return &PlacesRepository{db: db}
}

func (r *PlacesRepository) CreatePlace(userId int, place models.Place) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var placeId int
	insertPlaceQuery := fmt.Sprintf("INSERT INTO %s (name, description, address, rating) VALUES ($1, $2, $3, $4) RETURNING id", PlacesTable)
	row := tx.QueryRow(insertPlaceQuery, place.Name, place.Description, place.Address, place.Rating)
	if err := row.Scan(&placeId); err != nil {
		tx.Rollback()
		return 0, err
	}

	insertUserPlacesQuery := fmt.Sprintf("INSERT INTO %s (user_id, place_id) VALUES ($1, $2)", UserPlacesTable)
	_, err = tx.Exec(insertUserPlacesQuery, userId, placeId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return placeId, tx.Commit()
}
