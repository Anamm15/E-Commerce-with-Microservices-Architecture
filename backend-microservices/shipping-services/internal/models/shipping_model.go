package models

type Shipment struct {
	ID              uint64  `gorm:"primary_key"`
	Courier         string  `gorm:"not null"`
	DestinationCity string  `gorm:"not null"`
	ServiceType     string  `gorm:"not null"`
	Weight          int32   `gorm:"not null"`
	Price           float32 `gorm:"not null"`
	Status          string  `gorm:"not null"`
	TrackingNumber  string  `gorm:"not null"`
	EstimatedDays   int32   `gorm:"not null"`
	Address         Address `gorm:"foreignKey:ShipmentID"`
	OrderID         uint64  `gorm:"not null"`
}
