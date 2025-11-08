package repositories

import (
	"context"

	"product-services/internal/dto"
	"product-services/internal/models"

	"gorm.io/gorm"
)

type ImageRepository interface {
	GetImageByProductID(ctx context.Context, productID uint64) (dto.ImageResponseDTO, error)
	CreateImage(ctx context.Context, image *models.Image) (dto.ImageResponseDTO, error)
	DeleteImage(ctx context.Context, imageID uint64) (dto.ImageResponseDTO, error)
}

type imageRepository struct {
	db *gorm.DB
}

func NewImageRepository(db *gorm.DB) ImageRepository {
	return &imageRepository{db: db}
}

func (r *imageRepository) GetImageByProductID(ctx context.Context, productID uint64) (dto.ImageResponseDTO, error) {
	var image dto.ImageResponseDTO
	if err := r.db.WithContext(ctx).
		Where("product_id = ?", productID).
		Find(&image).Error; err != nil {
		return dto.ImageResponseDTO{}, err
	}
	return image, nil
}

func (r *imageRepository) CreateImage(ctx context.Context, image *models.Image) (dto.ImageResponseDTO, error) {
	if err := r.db.WithContext(ctx).
		Create(&image).Error; err != nil {
		return dto.ImageResponseDTO{}, err
	}
	return dto.ImageResponseDTO{ID: image.ID, URL: image.URL}, nil
}

func (r *imageRepository) DeleteImage(ctx context.Context, imageID uint64) (dto.ImageResponseDTO, error) {
	var image dto.ImageResponseDTO
	if err := r.db.WithContext(ctx).
		Where("id = ?", imageID).
		Delete(&image).Error; err != nil {
		return dto.ImageResponseDTO{}, err
	}
	return image, nil
}
