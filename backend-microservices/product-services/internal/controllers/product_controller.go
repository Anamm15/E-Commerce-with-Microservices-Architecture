package controllers

import (
	"context"

	"product-services/internal/constants"
	"product-services/internal/dto"
	"product-services/internal/services"
	"product-services/internal/utils"
	pb "product-services/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ProductServer struct {
	pb.UnimplementedProductServiceServer
	productService services.ProductService
}

func NewProductServer(productService services.ProductService) *ProductServer {
	return &ProductServer{
		productService: productService,
	}
}

func (s *ProductServer) GetAllProducts(ctx context.Context, req *emptypb.Empty) (*pb.ProductListResponse, error) {
	products, err := s.productService.GetAllProducts(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, constants.ErrProductServiceGet, err)
	}

	var res pb.ProductListResponse
	for _, product := range products {
		categories := utils.MapCategoryDTOResponseTogRPC(product)
		images := utils.MapImageDTOResponseTogRPC(product)
		res.Products = append(res.Products, &pb.ProductResponse{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			Rating:      product.Rating,
			OldPrice:    product.OldPrice,
			Category:    categories,
			ImageUrl:    images,
		})
	}

	return &res, nil
}

func (s *ProductServer) GetProductByID(ctx context.Context, req *pb.GetProductByIDRequest) (*pb.ProductResponse, error) {
	productID := req.GetId()

	product, err := s.productService.GetProductByID(ctx, productID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, constants.ErrProductServiceGetByID, err)
	}

	categories := utils.MapCategoryDTOResponseTogRPC(product)
	images := utils.MapImageDTOResponseTogRPC(product)

	return &pb.ProductResponse{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Rating:      product.Rating,
		OldPrice:    product.OldPrice,
		Category:    categories,
		ImageUrl:    images,
	}, nil
}

func (s *ProductServer) GetProductByCategoryID(ctx context.Context, req *pb.GetProductByCategoryRequest) (*pb.ProductListResponse, error) {
	categoryID := req.GetCategoryId()

	products, err := s.productService.GetProductByCategoryID(ctx, categoryID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, constants.ErrProductServiceGetByCategoryID, err)
	}

	var res pb.ProductListResponse
	for _, product := range products {
		categories := utils.MapCategoryDTOResponseTogRPC(product)
		images := utils.MapImageDTOResponseTogRPC(product)

		res.Products = append(res.Products, &pb.ProductResponse{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			Rating:      product.Rating,
			OldPrice:    product.OldPrice,
			Category:    categories,
			ImageUrl:    images,
		})
	}

	return &res, nil
}

func (s *ProductServer) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.ProductResponse, error) {
	dtoReq := dto.CreateProductRequestDTO{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		OldPrice:    req.OldPrice,
		Stock:       req.Stock,
		Category:    req.Category,
	}

	product, err := s.productService.CreateProduct(ctx, dtoReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, constants.ErrProductServiceCreate, err)
	}

	categories := utils.MapCategoryDTOResponseTogRPC(product)
	images := utils.MapImageDTOResponseTogRPC(product)

	return &pb.ProductResponse{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Rating:      product.Rating,
		OldPrice:    product.OldPrice,
		Category:    categories,
		ImageUrl:    images,
	}, nil
}

func (s *ProductServer) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.ProductResponse, error) {
	productID := req.GetId()

	dtoReq := dto.UpdateProductRequestDTO{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		OldPrice:    req.OldPrice,
		Stock:       req.Stock,
		Category:    req.Category,
	}

	product, err := s.productService.UpdateProduct(ctx, productID, dtoReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, constants.ErrProductServiceUpdate, err)
	}

	categories := utils.MapCategoryDTOResponseTogRPC(product)
	images := utils.MapImageDTOResponseTogRPC(product)

	return &pb.ProductResponse{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Rating:      product.Rating,
		OldPrice:    product.OldPrice,
		Category:    categories,
		ImageUrl:    images,
	}, nil
}

func (s *ProductServer) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*emptypb.Empty, error) {
	productID := req.GetId()

	err := s.productService.DeleteProduct(ctx, productID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, constants.ErrProductServiceDelete, err)
	}

	return &emptypb.Empty{}, nil
}
