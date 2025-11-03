package dto

import (
	"time"

	"product-services/models"
)

type ReviewResponseDTO struct {
	ID      uint        `json:"id"`
	Rating  int         `json:"rating"`
	Comment string      `json:"comment"`
	Date    time.Time   `json:"date"`
	User    models.User `json:"user"`
}

type CreateReviewRequestDTO struct {
	ProductID uint   `json:"product_id" validate:"required"`
	UserID    uint   `json:"user_id" validate:"required"`
	Rating    int    `json:"rating" validate:"required"`
	Comment   string `json:"comment" validate:"required"`
}

type UpdateReviewRequestDTO struct {
	ProductID uint   `json:"product_id" validate:"required"`
	UserID    uint   `json:"user_id" validate:"required"`
	Rating    int    `json:"rating" validate:"required"`
	Comment   string `json:"comment" validate:"required"`
}
