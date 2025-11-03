package dto

type StatusHistoryResponseDto struct {
	ID      uint   `json:"id"`
	OrderID uint   `json:"order_id"`
	Status  string `json:"status"`
	Date    string `json:"date"`
}

type CreateStatusHistoryRequestDto struct {
	OrderID uint   `json:"order_id"`
	Status  string `json:"status"`
}
