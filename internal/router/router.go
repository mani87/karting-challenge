package router

import (
	"github.com/gin-gonic/gin"
	"karting-challenge/internal/config"
	"karting-challenge/internal/handlers"
	"karting-challenge/internal/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/product", middleware.IsAuthentic(config.ApiKey), handlers.ListProducts)
		api.GET("/product/:productId", middleware.IsAuthentic(config.ApiKey), handlers.GetProduct)

		api.POST("/order", middleware.IsAuthentic(config.CreateOrderApiKey), handlers.Order)
	}

	return r
}
