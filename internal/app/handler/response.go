package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(ctx *gin.Context, statusCode int, message string) {
	log.Println(message)
	ctx.AbortWithStatusJSON(statusCode, ErrorResponse{Message: message})
}

type StatusResponse struct {
	Status string `json:"status"`
}
