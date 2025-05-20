package router

import (
	"github.com/gin-gonic/gin"
	"karting-challenge/handlers"
	"karting-challenge/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/product", middlewares.IsAuthentic(), handlers.ListProducts)
		api.GET("/product/:productId", handlers.GetProduct)
	}

	return r
}
