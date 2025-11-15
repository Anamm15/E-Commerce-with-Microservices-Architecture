package models

import "time"

type Order struct {
	ID                uint64      `gorm:"primaryKey"`
	Date              time.Time   `gorm:"not null, autoCreateTime"`
	Total             float32     `gorm:"not null"`
	Status            string      `gorm:"not null"`
	ShippingCost      float32     `gorm:"not null"`
	TrackingNumber    string      `gorm:"not null"`
	EstimatedDelivery int32       `gorm:"not null"`
	PaymentMethod     string      `gorm:"not null"`
	UserID            uint64      `gorm:"not null"`
	OrderItems        []OrderItem `gorm:"foreignKey:OrderID"`
	StatusHistory     []StatusHistory
	createdAt         time.Time
	updatedAt         time.Time
}
