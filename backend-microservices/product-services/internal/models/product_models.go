package models

import "time"

type Product struct {
	ID          uint64  `gorm:"primaryKey"`
	Name        string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	Price       float32 `gorm:"not null"`
	OldPrice    float32
	Stock       int32 `gorm:"not null"`
	Rating      float32
	IsNew       bool `gorm:"default:false"`
	Reviews     []Review
	ImageUrl    []Image    `gorm:"foreignKey:ProductID"`
	Category    []Category `gorm:"many2many:product_categories;"`
	CreatedAt   time.Time  `gorm:"not null default:current_timestamp"`
	UpdatedAt   time.Time
}
