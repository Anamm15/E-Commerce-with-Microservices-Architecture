package dto

type OrderItemResponseDto struct {
	ID        uint64
	ProductID uint64
	Quantity  int32
	Total     float32
}

type CreateOrderItemRequestDto struct {
	ProductID uint64  `json:"product_id"`
	Quantity  int32   `json:"quantity"`
	Total     float32 `json:"total"`
}
