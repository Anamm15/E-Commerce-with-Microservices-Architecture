package dto

type OrderItemResponseDto struct {
	ID        uint64
	OrderID   uint64
	ProductID uint64
	Name      string
	Quantity  int32
	Total     float32
	ImageUrl  string
}

type CreateOrderItemRequestDto struct {
	OrderID   uint64
	ProductID uint64
	Quantity  int32
	Total     float32
}
