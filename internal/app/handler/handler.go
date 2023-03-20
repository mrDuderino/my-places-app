package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mrDuderino/my-places-app/internal/app/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		places := api.Group("/places")
		{
			places.POST("/", h.createPlace)
			places.GET("/", h.getAllPlaces)
			places.GET("/:id", h.getPlaceById)
			places.GET("/name/:name", h.getPlaceByName)
			places.PUT("/:id", h.updatePlace)
			places.DELETE("/:id", h.deletePlace)

			dishes := places.Group(":id/dishes")
			{
				dishes.POST("/", h.createDish)
				dishes.GET("/", h.getAllPlaceDishes)
			}

			discounts := places.Group(":id/discounts")
			{
				discounts.POST("/", h.createDiscountCard)
				discounts.GET("/", h.getAllDiscountCards)
			}
		}

		dishes := api.Group("/dishes")
		{
			dishes.GET("/:id", h.getDishById)
			dishes.GET("/name/:name", h.getDishByName)
			dishes.PUT("/:id", h.updateDish)
			dishes.DELETE("/:id", h.deleteDish)
		}

		discounts := api.Group("/discounts")
		{
			discounts.GET("/:id", h.getDiscountCardById)
			discounts.PUT("/:id", h.updateDiscountCard)
			discounts.DELETE("/:id", h.deleteDiscountCard)
		}
	}
	
	return router
}
