package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"karting-challenge/internal/config"
	"karting-challenge/internal/handlers"
	"karting-challenge/internal/middleware"
	"time"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "x-api-key"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api")
	{
		api.GET("/product", middleware.IsAuthentic(config.ApiKey), handlers.ListProducts)
		api.GET("/product/:productId", middleware.IsAuthentic(config.ApiKey), handlers.GetProduct)

		api.POST("/order", middleware.IsAuthentic(config.CreateOrderApiKey), handlers.Order)
	}

	return r
}
