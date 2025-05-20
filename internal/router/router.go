package router

import (
	"github.com/gin-gonic/gin"
	"karting-challenge/internal/handlers"
	"karting-challenge/internal/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/product", middleware.IsAuthentic(), handlers.ListProducts)
		api.GET("/product/:productId", middleware.IsAuthentic(), handlers.GetProduct)
	}

	return r
}
