package models

import "time"

type Order struct {
	ID                uint      `gorm:"primaryKey"`
	Date              time.Time `gorm:"not null, autoCreateTime"`
	Total             float32   `gorm:"not null"`
	Status            string    `gorm:"not null"`
	ShippingCost      float32   `gorm:"not null"`
	TrackingNumber    string    `gorm:"not null"`
	EstimatedDelivery int       `gorm:"not null"`
	PaymentMethod     string    `gorm:"not null"`
	UserID            uint      `gorm:"not null"`
	OrderItems        []OrderItem
	StatusHistory     []StatusHistory
	createdAt         time.Time
	updatedAt         time.Time
}
