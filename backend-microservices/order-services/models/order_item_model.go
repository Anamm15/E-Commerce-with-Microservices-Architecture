package models

import "time"

type OrderItem struct {
	ID        uint    `gorm:"primaryKey"`
	ProductID uint    `gorm:"not null"`
	OrderID   uint    `gorm:"not null"`
	Quantity  int     `gorm:"not null"`
	Total     float32 `gorm:"not null"`
	createdAt time.Time
	updatedAt time.Time
}
