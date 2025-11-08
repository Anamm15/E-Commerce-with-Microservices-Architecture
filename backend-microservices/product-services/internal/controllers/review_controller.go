package controllers

import (
	"context"

	"product-services/internal/constants"
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
		return nil, status.Errorf(codes.NotFound, constants.ErrReviewServiceGet, err)
	}

	var res pb.ReviewListResponse
	for _, review := range reviews {
		var userReview pb.UserReview
		userReview.Id = review.User.ID
		userReview.FullName = review.User.FullName
		userReview.AvatarUrl = review.User.AvatarUrl

		res.Reviews = append(res.Reviews, &pb.ReviewResponse{
			Id:      review.ID,
			Rating:  review.Rating,
			Date:    timestamppb.New(review.Date),
			Comment: review.Comment,
			User:    &userReview,
		})
	}

	return &res, nil
}

func (s *ReviewServer) GetReviewByProductID(ctx context.Context, req *pb.GetReviewByProductRequest) (*pb.ReviewListResponse, error) {
	productID := req.GetProductId()

	reviews, err := s.reviewService.GetReviewProductID(ctx, productID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, constants.ErrReviewServiceGetByProductID, err)
	}

	var res pb.ReviewListResponse
	for _, review := range reviews {
		var userReview pb.UserReview
		userReview.Id = review.User.ID
		userReview.FullName = review.User.FullName
		userReview.AvatarUrl = review.User.AvatarUrl

		res.Reviews = append(res.Reviews, &pb.ReviewResponse{
			Id:      review.ID,
			Rating:  review.Rating,
			Date:    timestamppb.New(review.Date),
			Comment: review.Comment,
			User:    &userReview,
		})
	}

	return &res, nil
}

func (s *ReviewServer) CreateReview(ctx context.Context, req *pb.CreateReviewRequest) (*pb.ReviewResponse, error) {
	dtoReq := dto.CreateReviewRequestDTO{
		ProductID: req.GetProductId(),
		UserID:    req.GetUserId(),
		Rating:    req.GetRating(),
		Comment:   req.GetComment(),
	}

	review, err := s.reviewService.CreateReview(ctx, dtoReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, constants.ErrReviewServiceCreate, err)
	}

	return &pb.ReviewResponse{
		Id:      review.ID,
		Rating:  review.Rating,
		Date:    timestamppb.New(review.Date),
		Comment: review.Comment,
	}, nil
}

func (s *ReviewServer) UpdateReview(ctx context.Context, req *pb.UpdateReviewRequest) (*pb.ReviewResponse, error) {
	reviewID := req.GetReviewId()

	dtoReq := dto.UpdateReviewRequestDTO{
		ProductID: req.GetProductId(),
		UserID:    req.GetUserId(),
		Rating:    req.GetRating(),
		Comment:   req.GetComment(),
	}

	review, err := s.reviewService.UpdateReview(ctx, reviewID, dtoReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, constants.ErrReviewServiceUpdate, err)
	}

	return &pb.ReviewResponse{
		Id:      review.ID,
		Rating:  int32(review.Rating),
		Date:    timestamppb.New(review.Date),
		Comment: review.Comment,
	}, nil
}

func (s *ReviewServer) DeleteReview(ctx context.Context, req *pb.DeleteReviewRequest) (*emptypb.Empty, error) {
	reviewID := req.GetReviewId()
	userID := req.GetUserId()

	err := s.reviewService.DeleteReview(ctx, reviewID, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, constants.ErrReviewServiceDelete, err)
	}

	return &emptypb.Empty{}, nil
}
