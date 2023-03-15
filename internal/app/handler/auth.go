package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mrDuderino/my-places-app/internal/app/models"
	"net/http"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var input models.User
	err := ctx.BindJSON(&input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.CreateUser(input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

func (h *Handler) signIn(ctx *gin.Context) {

}
