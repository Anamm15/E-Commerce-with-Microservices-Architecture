package models

import "time"

type StatusHistory struct {
	ID      uint      `gorm:"primaryKey"`
	OrderID uint      `gorm:"not null"`
	Status  string    `gorm:"not null"`
	Date    time.Time `gorm:"autoCreateTime"`
}
