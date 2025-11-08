package models

import "time"

type Category struct {
	ID        uint64    `gorm:"primaryKey"`
	Name      string    `gorm:"not null;unique"`
	CreatedAt time.Time `gorm:"not null default:current_timestamp"`
	UpdatedAt time.Time `gorm:"not null default:current_timestamp"`
}
