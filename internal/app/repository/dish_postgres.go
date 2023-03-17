package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/mrDuderino/my-places-app/internal/app/models"
)

type DishRepository struct {
	db *sqlx.DB
}

func NewDishRepository(db *sqlx.DB) *DishRepository {
	return &DishRepository{db: db}
}

func (r *DishRepository) CreateDish(placeId int, dish models.Dish) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	dishQuery := fmt.Sprintf("INSERT INTO %s (name, description, rating) VALUES ($1, $2, $3) RETURNING id", DishesTable)
	row := tx.QueryRow(dishQuery, dish.Name, dish.Description, dish.Rating)
	var id int
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	placeDishesQuery := fmt.Sprintf("INSERT INTO %s (place_id, dish_id) VALUES ($1, $2)", PlaceDishesTable)
	_, err = tx.Exec(placeDishesQuery, placeId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}
