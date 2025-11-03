package dto

import (
	"order-services/models"
)

type OrderResponseDto struct {
	ID                uint    `json:"id"`
	Status            string  `json:"status"`
	Date              string  `json:"date"`
	Total             string  `json:"total"`
	ShippingCost      float32 `json:"shipping_cost"`
	TrackingNumber    string  `json:"tracking_number"`
	EstimatedDelivery int     `json:"estimated_delivery"`
	PaymentMethod     string  `json:"payment_method"`
	OrderItems        []models.OrderItem
}

type CreateOrderRequestDto struct {
	UserID            uint    `json:"user_id"`
	Total             float32 `json:"total"`
	ShippingCost      float32 `json:"shipping_cost"`
	TrackingNumber    string  `json:"tracking_number"`
	EstimatedDelivery int     `json:"estimated_delivery"`
	PaymentMethod     string  `json:"payment_method"`
}

type UpdateStatusOrderRequestDto struct {
	OrderID uint   `json:"order_id"`
	Status  string `json:"status"`
}
