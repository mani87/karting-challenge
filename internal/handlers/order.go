package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"karting-challenge/internal/models"
	"net/http"
)

func Order(c *gin.Context) {
	var req models.OrderReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid order"})
		return
	}

	// todo: validate promo code

	// Create Order
	order := models.Order{
		ID:        uuid.New().String(),
		Items:     req.Items,
		Discounts: 0.0,
		Total:     0.0,
		Products:  []models.Product{},
	}

	for _, item := range req.Items {
		for _, p := range models.Products {
			if p.ID == item.ProductID {
				order.Products = append(order.Products, p)
				order.Total += float64(item.Quantity) * p.Price
			}
		}
	}

	// todo: Take care of discount here

	c.JSON(http.StatusOK, order)
}
