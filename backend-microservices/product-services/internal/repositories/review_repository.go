package repositories

import (
	"context"

	"product-services/internal/dto"
	"product-services/internal/models"

	"gorm.io/gorm"
)

type ReviewRepository interface {
	GetAllReviews(ctx context.Context) ([]dto.ReviewResponseDTO, error)
	GetReviewByProductId(ctx context.Context, productId uint64) ([]dto.ReviewResponseDTO, error)
	CreateReview(ctx context.Context, review *models.Review) (dto.ReviewResponseDTO, error)
	UpdateReview(ctx context.Context, review *models.Review) (dto.ReviewResponseDTO, error)
	DeleteReview(ctx context.Context, reviewId uint64, userID uint64) error
}

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{db: db}
}

func (r *reviewRepository) GetAllReviews(ctx context.Context) ([]dto.ReviewResponseDTO, error) {
	var reviews []models.Review
	if err := r.db.WithContext(ctx).
		Find(&reviews).Error; err != nil {
		return nil, err
	}

	var reviewsDTOs []dto.ReviewResponseDTO
	for _, review := range reviews {
		reviewsDTOs = append(reviewsDTOs, dto.ReviewResponseDTO{
			ID:      review.ID,
			Rating:  review.Rating,
			Comment: review.Comment,
			Date:    review.Date,
		})
	}
	return reviewsDTOs, nil
}

func (r *reviewRepository) GetReviewByProductId(ctx context.Context, productId uint64) ([]dto.ReviewResponseDTO, error) {
	var reviews []models.Review
	if err := r.db.WithContext(ctx).
		Where("product_id = ?", productId).
		Find(&reviews).Error; err != nil {
		return nil, err
	}

	var reviewsDTOs []dto.ReviewResponseDTO
	for _, review := range reviews {
		reviewsDTOs = append(reviewsDTOs, dto.ReviewResponseDTO{
			ID:      review.ID,
			Rating:  review.Rating,
			Comment: review.Comment,
			Date:    review.Date,
		})
	}
	return reviewsDTOs, nil
}

func (r *reviewRepository) CreateReview(ctx context.Context, review *models.Review) (dto.ReviewResponseDTO, error) {
	if err := r.db.WithContext(ctx).
		Create(&review).Error; err != nil {
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
	if err := r.db.WithContext(ctx).
		Save(&review).Error; err != nil {
		return dto.ReviewResponseDTO{}, err
	}

	return dto.ReviewResponseDTO{
		ID:      review.ID,
		Rating:  review.Rating,
		Comment: review.Comment,
		Date:    review.Date,
	}, nil
}

func (r *reviewRepository) DeleteReview(ctx context.Context, reviewId uint64, userID uint64) error {
	if err := r.db.WithContext(ctx).
		Where("id = ? and user_id = ?", reviewId, userID).
		Delete(&models.Review{}).Error; err != nil {
		return err
	}
	return nil
}
