package models

import "time"

type User struct {
	ID          uint          `gorm:"primaryKey"`
	FullName    string        `gorm:"type:varchar(100)"`
	Username    string        `gorm:"unique;not null"`
	Email       string        `gorm:"unique;not null"`
	AvatarUrl   string        `gorm:"type:text"`
	Password    string        `gorm:"not null"`
	Role        string        `gorm:"not null; default:user"`
	IsConfirmed bool          `gorm:"default:false"`
	Addresses   []UserAddress `gorm:"foreignKey:UserID"`
	MemberSince time.Time     `gorm:"autoCreateTime"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
