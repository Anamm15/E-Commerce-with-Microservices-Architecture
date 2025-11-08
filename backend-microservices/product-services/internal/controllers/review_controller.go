package controllers

import (
	"context"
	"product-services/internal/dto"
	"product-services/internal/services"

	pb "product-services/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ReviewServer struct {
	pb.UnimplementedReviewServiceServer
	reviewService services.ReviewService
}

func NewReviewServer(reviewService services.ReviewService) *ReviewServer {
	return &ReviewServer{
		reviewService: reviewService,
	}
}

func (s *ReviewServer) GetAllReviews(ctx context.Context, req *emptypb.Empty) (*pb.ReviewListResponse, error) {
	reviews, err := s.reviewService.GetAllReviews(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to get reviews: %v", err)
	}

	var res pb.ReviewListResponse
	for _, review := range reviews {
		var userReview pb.UserReview
		userReview.Id = uint32(review.User.ID)
		userReview.FullName = review.User.FullName
		userReview.AvatarUrl = review.User.AvatarUrl

		res.Reviews = append(res.Reviews, &pb.ReviewResponse{
			Id:      uint32(review.ID),
			Rating:  int32(review.Rating),
			Date:    timestamppb.New(review.Date),
			Comment: review.Comment,
			User:    &userReview,
		})
	}

	return &res, nil
}

func (s *ReviewServer) GetReviewByProductID(ctx context.Context, req *pb.GetReviewByProductRequest) (*pb.ReviewListResponse, error) {
	productID := uint(req.GetProductId())

	reviews, err := s.reviewService.GetReviewProductID(ctx, productID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "reviews not found for product: %v", err)
	}

	var res pb.ReviewListResponse
	for _, review := range reviews {
		var userReview pb.UserReview
		userReview.Id = uint32(review.User.ID)
		userReview.FullName = review.User.FullName
		userReview.AvatarUrl = review.User.AvatarUrl

		res.Reviews = append(res.Reviews, &pb.ReviewResponse{
			Id:      uint32(review.ID),
			Rating:  int32(review.Rating),
			Date:    timestamppb.New(review.Date),
			Comment: review.Comment,
			User:    &userReview,
		})
	}

	return &res, nil
}

func (s *ReviewServer) CreateReview(ctx context.Context, req *pb.CreateReviewRequest) (*pb.ReviewResponse, error) {
	dtoReq := dto.CreateReviewRequestDTO{
		ProductID: uint(req.GetProductId()),
		UserID:    uint(req.GetUserId()),
		Rating:    int(req.GetRating()),
		Comment:   req.GetComment(),
	}

	review, err := s.reviewService.CreateReview(ctx, dtoReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create review: %v", err)
	}

	return &pb.ReviewResponse{
		Id:      uint32(review.ID),
		Rating:  int32(review.Rating),
		Date:    timestamppb.New(review.Date),
		Comment: review.Comment,
	}, nil
}

func (s *ReviewServer) UpdateReview(ctx context.Context, req *pb.UpdateReviewRequest) (*pb.ReviewResponse, error) {
	reviewID := uint(req.GetReviewId())

	dtoReq := dto.UpdateReviewRequestDTO{
		ProductID: uint(req.GetProductId()),
		UserID:    uint(req.GetUserId()),
		Rating:    int(req.GetRating()),
		Comment:   req.GetComment(),
	}

	review, err := s.reviewService.UpdateReview(ctx, reviewID, dtoReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update review: %v", err)
	}

	return &pb.ReviewResponse{
		Id:      uint32(review.ID),
		Rating:  int32(review.Rating),
		Date:    timestamppb.New(review.Date),
		Comment: review.Comment,
	}, nil
}

func (s *ReviewServer) DeleteReview(ctx context.Context, req *pb.DeleteReviewRequest) (*emptypb.Empty, error) {
	reviewID := uint(req.GetReviewId())

	err := s.reviewService.DeleteReview(ctx, reviewID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete review: %v", err)
	}

	return &emptypb.Empty{}, nil
}
