package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Order(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Order Placed"})
}
