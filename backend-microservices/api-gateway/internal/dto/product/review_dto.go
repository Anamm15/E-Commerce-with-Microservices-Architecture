package dto

import (
	"time"
)

type UserReviewResponse struct {
	ID        uint64 `json:"id"`
	FullName  string `json:"full_name"`
	AvatarUrl string `json:"avatar_url"`
}

type ReviewResponseDTO struct {
	ID      uint64             `json:"id"`
	Rating  int32              `json:"rating"`
	Comment string             `json:"comment"`
	Date    time.Time          `json:"date"`
	User    UserReviewResponse `json:"user"`
}

type CreateReviewRequestDTO struct {
	ProductID uint64 `json:"product_id" validate:"required"`
	Rating    int32  `json:"rating" validate:"required"`
	Comment   string `json:"comment" validate:"required"`
}

type UpdateReviewRequestDTO struct {
	ProductID uint64 `json:"product_id"`
	Rating    int32  `json:"rating"`
	Comment   string `json:"comment"`
}
