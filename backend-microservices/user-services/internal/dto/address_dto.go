package dto

type AddressResponseDTO struct {
	ID            uint64 `json:"id"`
	UserID        uint64 `json:"user_id"`
	Label         string `json:"label"`
	RecipientName string `json:"recipient_name"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	City          string `json:"city"`
	PostalCode    string `json:"postal_code"`
	IsDefault     bool   `json:"is_default"`
}

type CreateAddressRequestDTO struct {
	UserID        uint64 `json:"user_id"`
	Label         string `json:"label"`
	RecipientName string `json:"recipient_name"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	City          string `json:"city"`
	PostalCode    string `json:"postal_code"`
}

type UpdateAddressRequestDTO struct {
	ID            uint64 `json:"id"`
	Label         string `json:"label"`
	RecipientName string `json:"recipient_name"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	City          string `json:"city"`
	PostalCode    string `json:"postal_code"`
	IsDefault     bool   `json:"is_default"`
}
