package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mrDuderino/my-places-app/internal/app/models"
	"net/http"
)

func (h *Handler) createPlace(ctx *gin.Context) {
	userId, err := h.GetUserId(ctx)
	if err != nil {
		return
	}

	var input models.Place
	err = ctx.BindJSON(&input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Place.CreatePlace(userId, input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

type getAllPLacesResponse struct {
	Data []models.Place `json:"data"`
}

func (h *Handler) getAllPlaces(ctx *gin.Context) {
	userId, err := h.GetUserId(ctx)
	if err != nil {
		return
	}

	var places []models.Place
	places, err = h.services.Place.GetAllPlaces(userId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, getAllPLacesResponse{
		Data: places,
	})
}

func (h *Handler) getPlaceByName(ctx *gin.Context) {

}

func (h *Handler) updatePlace(ctx *gin.Context) {

}

func (h *Handler) deletePlace(ctx *gin.Context) {

}
