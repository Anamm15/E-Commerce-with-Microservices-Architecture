package models

import "time"

type User struct {
	ID   uint
	Name string
}

type Review struct {
	ID        uint      `gorm:"primaryKey"`
	ProductID uint      `gorm:"not null; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    uint      `gorm:"not null; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Rating    int       `gorm:"not null"`
	Comment   string    `gorm:"not null"`
	Date      time.Time `gorm:"not null default:current_timestamp"`
	User      User      `gorm:"idx:idx_review_user"`
	CreatedAt time.Time `gorm:"not null default:current_timestamp"`
	UpdatedAt time.Time
}
