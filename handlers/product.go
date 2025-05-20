package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Getting all products"})
}

func GetProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Getting product"})
}
