package dto

import (
	"mime/multipart"
)

type ImageResponseDTO struct {
	ID  uint   `json:"id"`
	URL string `json:"URL"`
}

type ProductResponseDTO struct {
	ID          uint                  `json:"id"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Price       float32               `json:"price"`
	OldPrice    float32               `json:"old_price"`
	Stock       int                   `json:"stock"`
	Rating      float32               `json:"rating"`
	IsNew       bool                  `json:"is_new"`
	Category    []CategoryResponseDTO `json:"category"`
	ImageUrl    []ImageResponseDTO    `json:"image_url"`
}

type CreateProductRequestDTO struct {
	Name        string                  `form:"name" validate:"required" binding:"required"`
	Description string                  `form:"description" validate:"required" binding:"required"`
	Category    []uint32                `form:"category" validate:"required" binding:"required"`
	Price       float32                 `form:"price" validate:"required" binding:"required"`
	OldPrice    float32                 `form:"old_price"`
	Stock       int                     `form:"stock" validate:"required" binding:"required"`
	IsNew       bool                    `form:"is_new"`
	Images      []*multipart.FileHeader `form:"images" validate:"required" binding:"required"`
}

type UpdateProductRequestDTO struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Category    []uint32 `json:"category"`
	Price       float32  `json:"price"`
	OldPrice    float32  `json:"old_price"`
	Stock       int      `json:"stock"`
	IsNew       bool     `json:"is_new"`
}
