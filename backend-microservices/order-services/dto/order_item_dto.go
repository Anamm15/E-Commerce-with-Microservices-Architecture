package dto

type OrderItemResponseDto struct {
	ID        uint
	OrderID   uint
	ProductID uint
	Quantity  int
	Total     float32
}

type CreateOrderItemRequestDto struct {
	OrderID   uint
	ProductID uint
	Quantity  int
	Total     float32
}
