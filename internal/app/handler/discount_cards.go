package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mrDuderino/my-places-app/internal/app/models"
	"net/http"
	"strconv"
)

func (h *Handler) createDiscountCard(ctx *gin.Context) {
	userId, err := h.GetUserId(ctx)
	if err != nil {
		return
	}

	placeId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return
	}

	var card models.DiscountCard
	if err := ctx.BindJSON(&card); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.DiscountCard.CreateDiscountCard(userId, placeId, card)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

type allDiscountCardsResponse struct {
	Data []models.DiscountCard `json:"data"`
}

func (h *Handler) getAllDiscountCards(ctx *gin.Context) {
	userId, err := h.GetUserId(ctx)
	if err != nil {
		return
	}

	placeId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return
	}

	var cards []models.DiscountCard
	cards, err = h.services.DiscountCard.GetAllDiscountCards(userId, placeId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, allDiscountCardsResponse{Data: cards})
}

func (h *Handler) getDiscountCardById(ctx *gin.Context) {
	userId, err := h.GetUserId(ctx)
	if err != nil {
		return
	}

	discountId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return
	}

	var card models.DiscountCard
	card, err = h.services.DiscountCard.GetById(userId, discountId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, card)
}

func (h *Handler) updateDiscountCard(ctx *gin.Context) {

}

func (h *Handler) deleteDiscountCard(ctx *gin.Context) {
	userId, err := h.GetUserId(ctx)
	if err != nil {
		return
	}

	discountId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return
	}

	err = h.services.DiscountCard.Delete(userId, discountId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}
