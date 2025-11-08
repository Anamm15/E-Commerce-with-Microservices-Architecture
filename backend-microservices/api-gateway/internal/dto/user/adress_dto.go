package user

type AddressResponseDTO struct {
	ID            uint   `json:"id"`
	UserID        uint   `json:"user_id"`
	Label         string `json:"label"`
	RecipientName string `json:"recipient_name"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	City          string `json:"city"`
	PostalCode    string `json:"postal_code"`
	IsDefault     bool   `json:"is_default"`
}

type CreateAddressRequestDTO struct {
	UserID        uint   `json:"user_id"`
	Label         string `json:"label" binding:"required"`
	RecipientName string `json:"recipient_name" binding:"required"`
	Phone         string `json:"phone" binding:"required"`
	Address       string `json:"address" binding:"required"`
	City          string `json:"city" binding:"required"`
	PostalCode    string `json:"postal_code" binding:"required"`
}

type UpdateAddressRequestDTO struct {
	ID            uint   `json:"id"`
	Label         string `json:"label"`
	RecipientName string `json:"recipient_name"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	City          string `json:"city"`
	PostalCode    string `json:"postal_code"`
	IsDefault     bool   `json:"is_default"`
	UserID        uint   `json:"user_id"`
}
