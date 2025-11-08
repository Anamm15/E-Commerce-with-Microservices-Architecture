package models

import "time"

type Review struct {
	ID        uint64    `gorm:"primaryKey"`
	ProductID uint64    `gorm:"not null; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    uint64    `gorm:"not null; idx:idx_review_user_id"`
	Rating    int32     `gorm:"not null"`
	Comment   string    `gorm:"not null"`
	Date      time.Time `gorm:"not null default:current_timestamp"`
	CreatedAt time.Time `gorm:"not null default:current_timestamp"`
	UpdatedAt time.Time
}
