package models

type Address struct {
	ID            uint64 `gorm:"primary_key"`
	RecipientName string `gorm:"not null"`
	Phone         string `gorm:"not null"`
	Street        string `gorm:"not null"`
	PostalCode    string `gorm:"not null"`
	ShipmentID    uint64 `gorm:"not null"`
}
