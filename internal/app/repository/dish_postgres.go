package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/mrDuderino/my-places-app/internal/app/models"
	"github.com/sirupsen/logrus"
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

func (r *DishRepository) GetAllDishes(userId, placeId int) ([]models.Dish, error) {
	query := fmt.Sprintf(`SELECT di.id, di.name, di.description, di.rating FROM %s di INNER JOIN %s pd ON di.id = pd.dish_id 
    INNER JOIN %s up ON pd.place_id = up.place_id WHERE up.user_id = $1 AND pd.place_id = $2`,
		DishesTable, PlaceDishesTable, UserPlacesTable)
	var dishes []models.Dish
	err := r.db.Select(&dishes, query, userId, placeId)
	return dishes, err
}

func (r *DishRepository) GetById(userId, dishId int) (models.Dish, error) {
	query := fmt.Sprintf(`SELECT di.id, di.name, di.description, di.rating FROM %s di INNER JOIN %s pd ON p.id=pd.dish_id 
    INNER JOIN %s up ON pd.place_id=up.place_id WHERE di.id=$1 AND up.user_id=$2`,
		DishesTable, PlaceDishesTable, UserPlacesTable)
	var dish models.Dish
	err := r.db.Get(&dish, query, dishId, userId)
	return dish, err
}

func (r *DishRepository) GetByName(userId int, dishName string) (models.Dish, error) {
	query := fmt.Sprintf("SELECT id FROM %s WHERE name=$1", DishesTable)
	var dishId int
	err := r.db.Get(&dishId, query, dishName)
	if err != nil {
		logrus.Warningf("error getting place id from place name for name=%s", dishName)
	}

	return r.GetById(userId, dishId)
}
