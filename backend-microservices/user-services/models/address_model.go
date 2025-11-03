package models

import "time"

type UserAddress struct {
	ID            uint   `gorm:"primaryKey"`
	UserID        uint   `gorm:"not null"`
	Label         string `gorm:"type:varchar(50)"`
	IsDefault     bool   `gorm:"default:false"`
	RecipientName string `gorm:"type:varchar(100)"`
	Phone         string `gorm:"type:varchar(20)"`
	Address       string `gorm:"type:varchar(100)"`
	City          string `gorm:"type:varchar(100)"`
	PostalCode    string `gorm:"type:varchar(20)"`
	User          User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
