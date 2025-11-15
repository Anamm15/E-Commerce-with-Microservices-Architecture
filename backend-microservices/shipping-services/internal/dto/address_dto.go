package dto

type AddressResponseDTO struct {
	ID            uint64 `json:"id"`
	RecipientName string `json:"recipient_name"`
	Phone         string `json:"phone"`
	Street        string `json:"street"`
	PostalCode    string `json:"postal_code"`
}
