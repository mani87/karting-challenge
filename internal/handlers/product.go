package handlers

import (
	"github.com/gin-gonic/gin"
	"karting-challenge/internal/services"
	"net/http"
)

func ListProducts(c *gin.Context) {
	c.JSON(http.StatusOK, services.ListProducts())
}

func GetProduct(c *gin.Context) {
	id := c.Param("productId")
	product, exists := services.GetProductByID(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}
