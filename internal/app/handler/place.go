package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createPlace(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"user_id": userId})
}

func (h *Handler) getAllPlaces(ctx *gin.Context) {

}

func (h *Handler) getPlaceByName(ctx *gin.Context) {

}

func (h *Handler) updatePlace(ctx *gin.Context) {

}

func (h *Handler) deletePlace(ctx *gin.Context) {

}
