package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/mrDuderino/my-places-app/internal/app/models"
	"strings"
)

type DiscountCardRepository struct {
	db *sqlx.DB
}

func NewDiscountCardRepository(db *sqlx.DB) *DiscountCardRepository {
	return &DiscountCardRepository{db: db}
}

func (r *DiscountCardRepository) CreateDiscountCard(placeId int, card models.DiscountCard) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf(`INSERT INTO %s (number, description, valid_from, valid_to) 
		VALUES ($1, $2, $3, $4) RETURNING id`, DiscountCardsTable)
	row := tx.QueryRow(query, card.Number, card.Description, card.ValidFrom, card.ValidTo)

	var cardId int
	if err := row.Scan(&cardId); err != nil {
		tx.Rollback()
		return 0, err
	}

	query = fmt.Sprintf("INSERT INTO %s (place_id, discount_card_id) VALUES ($1, $2)", PlaceDiscountCardsTable)
	if _, err := tx.Exec(query, placeId, cardId); err != nil {
		tx.Rollback()
		return 0, err
	}

	return cardId, tx.Commit()
}

func (r *DiscountCardRepository) GetAllDiscountCards(userId, placeId int) ([]models.DiscountCard, error) {
	query := fmt.Sprintf(`SELECT dc.id, dc.number, dc.description, dc.valid_from, dc.valid_to FROM %s dc 
    	INNER JOIN %s pd ON dc.id=pd.discount_card_id INNER JOIN %s up ON pd.place_id=up.place_id WHERE up.user_id=$1 AND pd.place_id=$2`,
		DiscountCardsTable, PlaceDiscountCardsTable, UserPlacesTable)

	var cards []models.DiscountCard
	err := r.db.Select(&cards, query, userId, placeId)

	return cards, err
}

func (r *DiscountCardRepository) GetById(userId, discountId int) (models.DiscountCard, error) {
	query := fmt.Sprintf(`SELECT dc.id, dc.number, dc.description, dc.valid_from, dc.valid_to FROM %s dc 
    	INNER JOIN %s pd ON dc.id=pd.discount_card_id INNER JOIN %s up ON pd.place_id=up.place_id WHERE up.user_id=$1 AND dc.id=$2`,
		DiscountCardsTable, PlaceDiscountCardsTable, UserPlacesTable)

	var card models.DiscountCard
	err := r.db.Get(&card, query, userId, discountId)

	return card, err
}

func (r *DiscountCardRepository) Delete(userId, discountId int) error {
	query := fmt.Sprintf(`DELETE FROM %s dc USING %s pd, %s up 
       	WHERE dc.id=pd.discount_card_id AND pd.place_id=up.place_id AND up.user_id=$1 AND dc.id=$2`,
		DiscountCardsTable, PlaceDiscountCardsTable, UserPlacesTable)
	_, err := r.db.Exec(query, userId, discountId)

	return err
}

func (r *DiscountCardRepository) Update(userId int, discountId int, input models.UpdateDiscountCardInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Number != nil {
		setValues = append(setValues, fmt.Sprintf("number=$%d", argId))
		args = append(args, *input.Number)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.ValidFrom != nil {
		setValues = append(setValues, fmt.Sprintf("valid_from=$%d", argId))
		args = append(args, *input.ValidFrom)
		argId++
	}

	if input.ValidTo != nil {
		setValues = append(setValues, fmt.Sprintf("valid_to=$%d", argId))
		args = append(args, *input.ValidTo)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf(`UPDATE %s dc SET %s FROM %s pd, %s up 
       	WHERE dc.id=pd.discount_card_id AND pd.place_id=up.place_id AND up.user_id=$%d AND dc.id=$%d`,
		DiscountCardsTable, setQuery, PlaceDiscountCardsTable, UserPlacesTable, argId, argId+1)
	args = append(args, userId, discountId)
	_, err := r.db.Exec(query, args...)

	return err
}
