package models

import "time"

type StatusHistory struct {
	ID      uint64    `gorm:"primaryKey"`
	OrderID uint64    `gorm:"not null"`
	Status  string    `gorm:"not null"`
	Date    time.Time `gorm:"autoCreateTime"`
}
