package services

import (
	"context"

	"product-services/dto"
	"product-services/models"
	"product-services/repositories"
)

type ReviewService interface {
	GetAllReviews(ctx context.Context) ([]dto.ReviewResponseDTO, error)
	GetReviewProductID(ctx context.Context, productID uint) (dto.ReviewResponseDTO, error)
	CreateReview(ctx context.Context, reviewRequest dto.CreateReviewRequestDTO) (dto.ReviewResponseDTO, error)
	UpdateReview(ctx context.Context, reviewID uint, reviewRequest dto.UpdateReviewRequestDTO) (dto.ReviewResponseDTO, error)
	DeleteReview(ctx context.Context, reviewID uint) error
}

type reviewService struct {
	reviewRepository repositories.ReviewRepository
}

func NewReviewService(reviewRepository repositories.ReviewRepository) ReviewService {
	return &reviewService{reviewRepository: reviewRepository}
}

func (s *reviewService) GetAllReviews(ctx context.Context) ([]dto.ReviewResponseDTO, error) {
	return s.reviewRepository.GetAllReviews(ctx)
}

func (s *reviewService) GetReviewProductID(ctx context.Context, productID uint) (dto.ReviewResponseDTO, error) {
	return s.reviewRepository.GetReviewByProductId(ctx, productID)
}

func (s *reviewService) CreateReview(ctx context.Context, reviewRequest dto.CreateReviewRequestDTO) (dto.ReviewResponseDTO, error) {
	review := models.Review{
		ProductID: reviewRequest.ProductID,
		Rating:    reviewRequest.Rating,
		Comment:   reviewRequest.Comment,
	}
	return s.reviewRepository.CreateReview(ctx, &review)
}

func (s *reviewService) UpdateReview(ctx context.Context, reviewID uint, reviewRequest dto.UpdateReviewRequestDTO) (dto.ReviewResponseDTO, error) {
	review := models.Review{
		ID:        reviewID,
		ProductID: reviewRequest.ProductID,
		Rating:    reviewRequest.Rating,
		Comment:   reviewRequest.Comment,
	}
	return s.reviewRepository.UpdateReview(ctx, &review)
}

func (s *reviewService) DeleteReview(ctx context.Context, reviewID uint) error {
	return s.reviewRepository.DeleteReview(ctx, reviewID)
}
