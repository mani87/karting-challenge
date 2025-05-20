package router

import (
	"github.com/gin-gonic/gin"
	"karting-challenge/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/product", handlers.ListProducts)
		api.GET("/product/:productId", handlers.GetProduct)
	}

	return r
}
