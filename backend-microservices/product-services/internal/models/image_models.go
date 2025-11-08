package models

import "time"

type Image struct {
	ID        uint64    `gorm:"primaryKey"`
	ProductID uint64    `gorm:"not null; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	URL       string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null default:current_timestamp"`
	UpdatedAt time.Time `gorm:"not null default:current_timestamp"`
}
