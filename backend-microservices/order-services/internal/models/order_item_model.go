package models

import "time"

type OrderItem struct {
	ID        uint64  `gorm:"primaryKey"`
	ProductID uint64  `gorm:"not null"`
	OrderID   uint64  `gorm:"not null"`
	Quantity  int32   `gorm:"not null"`
	Total     float32 `gorm:"not null"`
	createdAt time.Time
	updatedAt time.Time
}
