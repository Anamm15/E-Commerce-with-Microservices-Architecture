package dto

type OrderResponseDto struct {
	ID                uint64                 `json:"id"`
	Status            string                 `json:"status"`
	Date              string                 `json:"date"`
	Total             float32                `json:"total"`
	ShippingCost      float32                `json:"shipping_cost"`
	TrackingNumber    string                 `json:"tracking_number"`
	EstimatedDelivery int32                  `json:"estimated_delivery"`
	PaymentMethod     string                 `json:"payment_method"`
	OrderItems        []OrderItemResponseDto `json:"order_items"`
}

type CreateOrderRequestDto struct {
	UserID        uint64                      `json:"user_id"`
	Total         float32                     `json:"total" binding:"required"`
	PaymentMethod string                      `json:"payment_method" binding:"required"`
	Item          []CreateOrderItemRequestDto `json:"item"`
}

type UpdateStatusOrderRequestDto struct {
	OrderID uint64 `json:"order_id" binding:"required"`
	Status  string `json:"status" binding:"required"`
}
