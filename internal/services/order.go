package services

import (
	"github.com/google/uuid"
	"karting-challenge/internal/models"
)

func PlaceOrder(req models.OrderReq) models.Order {
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

	if len(req.CouponCode) > 0 {
		order.Discounts = order.Total * 0.80 // 20% discount if its valid coupon
	}

	return order
}
