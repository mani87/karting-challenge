package models

type Image struct {
	Thumbnail string `json:"thumbnail"`
	Mobile    string `json:"mobile"`
	Tablet    string `json:"tablet"`
	Desktop   string `json:"desktop"`
}

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
	Image    Image   `json:"image"`
}

type OrderReq struct {
	CouponCode string         `json:"couponCode"`
	Items      []OrderItemReq `json:"items" binding:"required"`
}

type OrderItemReq struct {
	ProductID string `json:"productId" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
}

type Order struct {
	ID        string         `json:"id"`
	Total     float64        `json:"total"`
	Discounts float64        `json:"discounts"`
	Items     []OrderItemReq `json:"items"`
	Products  []Product      `json:"products"`
}
