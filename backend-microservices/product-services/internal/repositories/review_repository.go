package repositories

import (
	"context"

	"product-services/internal/dto"
	"product-services/internal/models"

	"gorm.io/gorm"
)

type ReviewRepository interface {
	GetAllReviews(ctx context.Context) ([]dto.ReviewResponseDTO, error)
	GetReviewByProductId(ctx context.Context, productId uint) ([]dto.ReviewResponseDTO, error)
	CreateReview(ctx context.Context, review *models.Review) (dto.ReviewResponseDTO, error)
	UpdateReview(ctx context.Context, review *models.Review) (dto.ReviewResponseDTO, error)
	DeleteReview(ctx context.Context, reviewId uint) error
}

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{db: db}
}

func (r *reviewRepository) GetAllReviews(ctx context.Context) ([]dto.ReviewResponseDTO, error) {
	var reviews []dto.ReviewResponseDTO
	if err := r.db.Find(&reviews).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}

func (r *reviewRepository) GetReviewByProductId(ctx context.Context, productId uint) ([]dto.ReviewResponseDTO, error) {
	var review []dto.ReviewResponseDTO
	if err := r.db.Where("product_id = ?", productId).Find(&review).Error; err != nil {
		return nil, err
	}
	return review, nil
}

func (r *reviewRepository) CreateReview(ctx context.Context, review *models.Review) (dto.ReviewResponseDTO, error) {
	if err := r.db.Create(&review).Error; err != nil {
		return dto.ReviewResponseDTO{}, err
	}
	return dto.ReviewResponseDTO{
		ID:      review.ID,
		Rating:  review.Rating,
		Comment: review.Comment,
		Date:    review.Date,
	}, nil
}

func (r *reviewRepository) UpdateReview(ctx context.Context, review *models.Review) (dto.ReviewResponseDTO, error) {
	if err := r.db.Save(&review).Error; err != nil {
		return dto.ReviewResponseDTO{}, err
	}
	return dto.ReviewResponseDTO{
		ID:      review.ID,
		Rating:  review.Rating,
		Comment: review.Comment,
		Date:    review.Date,
	}, nil
}

func (r *reviewRepository) DeleteReview(ctx context.Context, reviewId uint) error {
	if err := r.db.Where("id = ?", reviewId).Delete(&models.Review{}).Error; err != nil {
		return err
	}
	return nil
}
