package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID          uint64        `gorm:"primaryKey"`
	FullName    string        `gorm:"type:varchar(100)"`
	Username    string        `gorm:"unique;not null"`
	Email       string        `gorm:"unique;not null"`
	AvatarUrl   string        `gorm:"type:text"`
	Password    string        `gorm:"not null"`
	Role        string        `gorm:"not null;default:user"`
	IsConfirmed bool          `gorm:"default:false"`
	Addresses   []UserAddress `gorm:"foreignKey:UserID"`
	MemberSince time.Time     `gorm:"autoCreateTime"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if len(u.Password) > 0 {
		hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashed)
	}
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed("Password") && len(u.Password) > 0 {
		hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashed)
	}
	return nil
}
