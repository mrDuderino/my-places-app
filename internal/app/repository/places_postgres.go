package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/mrDuderino/my-places-app/internal/app/models"
	"github.com/sirupsen/logrus"
	"strings"
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

func (r *PlacesRepository) GetAllPlaces(userId int) ([]models.Place, error) {
	query := fmt.Sprintf(`SELECT p.id, p.name, p.description, p.address, p.rating FROM %s p 
                            INNER JOIN %s up ON p.id = up.place_id WHERE up.user_id = $1`, PlacesTable, UserPlacesTable)
	var places []models.Place
	err := r.db.Select(&places, query, userId)
	return places, err
}

func (r *PlacesRepository) GetById(userId, placeId int) (models.Place, error) {
	query := fmt.Sprintf(`SELECT p.id, p.name, p.description, p.address, p.rating FROM %s p 
                            INNER JOIN %s up ON p.id = up.place_id WHERE up.user_id = $1 AND p.id = $2`, PlacesTable, UserPlacesTable)
	var place models.Place
	err := r.db.Get(&place, query, userId, placeId)
	return place, err
}

func (r *PlacesRepository) GetByName(userId int, placeName string) (models.Place, error) {
	query := fmt.Sprintf("SELECT id FROM %s WHERE name = $1", PlacesTable)
	var placeId int
	err := r.db.Get(&placeId, query, placeName)
	if err != nil {
		logrus.Warningf("error getting place id from place name for name=%s", placeName)
	}

	return r.GetById(userId, placeId)
}

func (r *PlacesRepository) Delete(userId, placeId int) error {
	query := fmt.Sprintf("DELETE FROM %s p USING %s up WHERE p.id = up.place_id AND p.id = $1 AND up.user_id = $2",
		PlacesTable, UserPlacesTable)
	_, err := r.db.Exec(query, placeId, userId)
	return err
}

func (r *PlacesRepository) Update(userId int, placeId int, input models.UpdatePlaceInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Address != nil {
		setValues = append(setValues, fmt.Sprintf("address=$%d", argId))
		args = append(args, *input.Address)
		argId++
	}

	if input.Rating != nil {
		setValues = append(setValues, fmt.Sprintf("rating=$%d", argId))
		args = append(args, *input.Rating)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s p SET %s FROM %s up WHERE p.id = up.place_id AND p.id = $%d AND up.user_id = $%d",
		PlacesTable, setQuery, UserPlacesTable, argId, argId+1)

	args = append(args, placeId, userId)
	_, err := r.db.Exec(query, args...)
	return err
}
