package handlers

import (
	"github.com/gin-gonic/gin"
	"karting-challenge/internal/models"
	"karting-challenge/internal/services"
	"karting-challenge/internal/utils"
	"net/http"
)

func Order(c *gin.Context) {
	var req models.OrderReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid order"})
		return
	}

	valid := utils.ValidatePromoCode(req.CouponCode)
	if !valid {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid promotion code"})
		return
	}

	order := services.PlaceOrder(req)
	c.JSON(http.StatusOK, order)
}
