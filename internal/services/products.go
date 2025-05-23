package services

import (
	"karting-challenge/internal/models"
)

func ListProducts() []models.Product {
	return models.Products
}

func GetProductByID(id string) (models.Product, bool) {
	for _, p := range models.Products {
		if p.ID == id {
			return p, true
		}
	}
	return models.Product{}, false
}
