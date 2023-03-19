package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mrDuderino/my-places-app/internal/app/models"
	"net/http"
	"strconv"
)

func (h *Handler) createDish(ctx *gin.Context) {
	userId, err := h.GetUserId(ctx)
	if err != nil {
		return
	}

	placeId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return
	}

	var dish models.Dish
	err = ctx.BindJSON(&dish)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	id, err := h.services.Dish.CreateDish(userId, placeId, dish)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

type allDishesResponse struct {
	Data []models.Dish `json:"data"`
}

func (h *Handler) getAllPlaceDishes(ctx *gin.Context) {
	userId, err := h.GetUserId(ctx)
	if err != nil {
		return
	}

	placeId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return
	}

	var dishes []models.Dish
	dishes, err = h.services.Dish.GetAllDishes(userId, placeId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, allDishesResponse{Data: dishes})
}

func (h *Handler) getDishById(ctx *gin.Context) {
	userId, err := h.GetUserId(ctx)
	if err != nil {
		return
	}

	dishId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "invalid dish id")
		return
	}

	var dish models.Dish
	dish, err = h.services.Dish.GetById(userId, dishId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, dish)
}

func (h *Handler) getDishByName(ctx *gin.Context) {

}

func (h *Handler) updateDish(ctx *gin.Context) {

}

func (h *Handler) deleteDish(ctx *gin.Context) {

}
