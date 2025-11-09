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
)

type CategoryServer struct {
	pb.UnimplementedCategoryServiceServer
	categoryService services.CategoryService
}

func NewCategoryServer(categoryService services.CategoryService) *CategoryServer {
	return &CategoryServer{
		categoryService: categoryService,
	}
}

func (s *CategoryServer) GetAllCategories(ctx context.Context, req *emptypb.Empty) (*pb.CategoryListResponse, error) {
	categories, err := s.categoryService.GetAllCategories(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, constants.ErrCategoryServiceGet, err)
	}

	var pbCategories []*pb.CategoryResponse
	for _, category := range categories {
		pbCategories = append(pbCategories, &pb.CategoryResponse{
			Id:   category.ID,
			Name: category.Name,
		})
	}

	return &pb.CategoryListResponse{
		Categories: pbCategories,
	}, nil
}

func (s *CategoryServer) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	dtoReq := dto.CreateCategoryRequestDTO{
		Name: req.GetName(),
	}

	category, err := s.categoryService.CreateCategory(ctx, dtoReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, constants.ErrCategoryServiceCreate, err)
	}

	pbCategory := &pb.CategoryResponse{
		Id:   category.ID,
		Name: category.Name,
	}

	return pbCategory, nil
}

func (s *CategoryServer) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryRequest) (*pb.CategoryResponse, error) {
	categoryID := req.GetId()
	dtoReq := dto.UpdateCategoryRequestDTO{
		Name: req.GetName(),
	}

	category, err := s.categoryService.UpdateCategory(ctx, categoryID, dtoReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, constants.ErrCategoryServiceUpdate, err)
	}

	pbCategory := &pb.CategoryResponse{
		Id:   category.ID,
		Name: category.Name,
	}

	return pbCategory, nil
}

func (s *CategoryServer) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryRequest) (*emptypb.Empty, error) {
	categoryID := req.GetId()

	err := s.categoryService.DeleteCategory(ctx, categoryID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, constants.ErrCategoryServiceDelete, err)
	}

	return &emptypb.Empty{}, nil
}
