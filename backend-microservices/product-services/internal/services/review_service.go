package services

import (
	"context"

	"product-services/internal/dto"
	"product-services/internal/models"
	"product-services/internal/repositories"
	pbUser "product-services/pb/user"
)

type ReviewService interface {
	GetAllReviews(ctx context.Context) ([]dto.ReviewResponseDTO, error)
	GetReviewProductID(ctx context.Context, productID uint) ([]dto.ReviewResponseDTO, error)
	CreateReview(ctx context.Context, reviewRequest dto.CreateReviewRequestDTO) (dto.ReviewResponseDTO, error)
	UpdateReview(ctx context.Context, reviewID uint, reviewRequest dto.UpdateReviewRequestDTO) (dto.ReviewResponseDTO, error)
	DeleteReview(ctx context.Context, reviewID uint) error
}

type reviewService struct {
	reviewRepository repositories.ReviewRepository
	userClient       pbUser.UserServiceClient
}

func NewReviewService(
	reviewRepository repositories.ReviewRepository,
	userClient pbUser.UserServiceClient,
) ReviewService {
	return &reviewService{
		reviewRepository: reviewRepository,
		userClient:       userClient,
	}
}

func (s *reviewService) GetAllReviews(ctx context.Context) ([]dto.ReviewResponseDTO, error) {
	reviews, err := s.reviewRepository.GetAllReviews(ctx)
	if err != nil {
		return nil, err
	}

	for i, r := range reviews {
		userResp, err := s.userClient.GetUserByID(ctx, &pbUser.GetUserRequest{Id: uint64(r.ID)})
		if err == nil {
			reviews[i].User.ID = uint(userResp.Id)
			reviews[i].User.FullName = userResp.FullName
			reviews[i].User.AvatarUrl = userResp.AvatarUrl
		}
	}

	return reviews, nil
}

func (s *reviewService) GetReviewProductID(ctx context.Context, productID uint) ([]dto.ReviewResponseDTO, error) {
	reviews, err := s.reviewRepository.GetReviewByProductId(ctx, productID)
	if err != nil {
		return nil, err
	}

	for i, r := range reviews {
		userResp, err := s.userClient.GetUserByID(ctx, &pbUser.GetUserRequest{Id: uint64(r.ID)})
		if err == nil {
			reviews[i].User.ID = uint(userResp.Id)
			reviews[i].User.FullName = userResp.FullName
			reviews[i].User.AvatarUrl = userResp.AvatarUrl
		}
	}

	return reviews, nil
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
