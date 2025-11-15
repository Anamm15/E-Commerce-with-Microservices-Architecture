package shippping

type ShipmentResponseDTO struct {
	ID              uint64             `json:"id"`
	Courier         string             `json:"courier"`
	DestinationCity string             `json:"destination_city"`
	ServiceType     string             `json:"service_type"`
	Weight          int32              `json:"weight"`
	Price           float32            `json:"price"`
	Status          string             `json:"status"`
	Address         AddressResponseDTO `json:"address"`
	TrackingNumber  string             `json:"tracking_number"`
	EstimatedDays   int32              `json:"estimated_days"`
}

type CalculateCostResponseDTO struct {
	Courier       string  `json:"courier"`
	ServiceType   string  `json:"service_type"`
	Price         float32 `json:"price"`
	EstimatedDays int32   `json:"estimated_days"`
}

type CalculateCostRequestDTO struct {
	OriginCity      string `json:"origin_city" binding:"required"`
	DestinationCity string `json:"destination_city" binding:"required"`
	PostalCode      string `json:"postal_code" binding:"required"`
	Weight          int32  `json:"weight" binding:"required"`
}
