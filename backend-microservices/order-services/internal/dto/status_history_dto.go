package dto

type StatusHistoryResponseDto struct {
	ID      uint64 `json:"id"`
	OrderID uint64 `json:"order_id"`
	Status  string `json:"status"`
	Date    string `json:"date"`
}

type CreateStatusHistoryRequestDto struct {
	OrderID uint64 `json:"order_id"`
	Status  string `json:"status"`
}
