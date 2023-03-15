package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mrDuderino/my-places-app/internal/app/models"
	"log"
	"net/http"
	"strconv"
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

func (h *Handler) getPlaceById(ctx *gin.Context) {
	userId, err := h.GetUserId(ctx)
	if err != nil {
		return
	}

	placeId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "invalid place id")
		return
	}

	var place models.Place
	place, err = h.services.Place.GetById(userId, placeId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, place)
}

func (h *Handler) getPlaceByName(ctx *gin.Context) {
	userId, err := h.GetUserId(ctx)
	if err != nil {
		return
	}

	placeName := ctx.Param("name")

	var place models.Place
	place, err = h.services.Place.GetByName(userId, placeName)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, place)
}

func (h *Handler) updatePlace(ctx *gin.Context) {
	userId, err := h.GetUserId(ctx)
	if err != nil {
		return
	}

	placeId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "invalid place id")
		return
	}

	var input models.UpdatePlaceInput
	err = ctx.BindJSON(&input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	log.Printf("%#+v", input)

	err = h.services.Place.Update(userId, placeId, input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}

func (h *Handler) deletePlace(ctx *gin.Context) {
	userId, err := h.GetUserId(ctx)
	if err != nil {
		return
	}

	placeId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "invalid place id")
		return
	}

	err = h.services.Place.Delete(userId, placeId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}
