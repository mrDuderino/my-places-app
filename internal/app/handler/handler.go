package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		places := api.Group("/places")
		{
			places.POST("/", h.createPlace)
			places.GET("/", h.getAllPlaces)
			places.GET("/:name", h.getPlaceByName)
			places.PUT("/:name", h.updatePlace)
			places.DELETE("/:name", h.deletePlace)
		}
	}
	return router
}
