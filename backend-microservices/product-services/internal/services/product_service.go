package services

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"

	"product-services/internal/dto"
	"product-services/internal/models"
	"product-services/internal/repositories"
	"product-services/internal/storages"
)

type ProductService interface {
	GetAllProducts(ctx context.Context) ([]dto.ProductResponseDTO, error)
	GetProductByID(ctx context.Context, productID uint64) (dto.ProductResponseDTO, error)
	GetProductByCategoryID(ctx context.Context, categoryID uint64) ([]dto.ProductResponseDTO, error)
	CreateProduct(ctx context.Context, productRequest dto.CreateProductRequestDTO) (dto.ProductResponseDTO, error)
	UpdateProduct(ctx context.Context, productID uint64, productRequest dto.UpdateProductRequestDTO) (dto.ProductResponseDTO, error)
	DeleteProduct(ctx context.Context, productID uint64) error
}

type productService struct {
	categoryRepository repositories.CategoryRepository
	productRepository  repositories.ProductRepository
	imageRepository    repositories.ImageRepository
	storage            *storages.FirebaseStorage
}

func NewProductService(
	categoryRepository repositories.CategoryRepository,
	productRepository repositories.ProductRepository,
	imageRepository repositories.ImageRepository,
	storage *storages.FirebaseStorage,
) ProductService {
	return &productService{
		categoryRepository: categoryRepository,
		productRepository:  productRepository,
		imageRepository:    imageRepository,
		storage:            storage,
	}
}

func (s *productService) GetAllProducts(ctx context.Context) ([]dto.ProductResponseDTO, error) {
	return s.productRepository.GetAllProducts(ctx)
}

func (s *productService) GetProductByID(ctx context.Context, productID uint64) (dto.ProductResponseDTO, error) {
	return s.productRepository.GetProductById(ctx, productID)
}

func (s *productService) GetProductByCategoryID(ctx context.Context, categoryID uint64) ([]dto.ProductResponseDTO, error) {
	return s.productRepository.GetProductsByCategory(ctx, categoryID)
}

func (s *productService) CreateProduct(ctx context.Context, productRequest dto.CreateProductRequestDTO) (dto.ProductResponseDTO, error) {
	product := models.Product{
		Name:        productRequest.Name,
		Description: productRequest.Description,
		Price:       productRequest.Price,
		OldPrice:    productRequest.OldPrice,
		Stock:       productRequest.Stock,
		IsNew:       productRequest.IsNew,
	}

	createdProduct, err := s.productRepository.CreateProduct(ctx, &product)
	if err != nil {
		return dto.ProductResponseDTO{}, fmt.Errorf("gagal membuat produk: %w", err)
	}
	if createdProduct.ID == 0 {
		return dto.ProductResponseDTO{}, fmt.Errorf("gagal membuat produk, ID tidak valid")
	}

	if len(productRequest.Category) > 0 {
		categories, err := s.categoryRepository.AddProductCategory(ctx, createdProduct.ID, productRequest.Category)
		if err != nil {
			return dto.ProductResponseDTO{}, fmt.Errorf("gagal menambahkan kategori produk: %w", err)
		}
		createdProduct.Category = categories
	}

	if len(productRequest.Images) > 0 {
		for i, imgBytes := range productRequest.Images {
			if len(imgBytes) == 0 {
				continue
			}

			path := fmt.Sprintf("e-commerce/products/%d/%d_%d.jpg", createdProduct.ID, time.Now().UnixNano(), i)
			contentType := http.DetectContentType(imgBytes[:512])

			reader := bytes.NewReader(imgBytes)
			url, err := s.storage.UploadFile(ctx, path, reader, contentType)
			if err != nil {
				return dto.ProductResponseDTO{}, fmt.Errorf("gagal upload file: %w", err)
			}

			image := models.Image{
				ProductID: createdProduct.ID,
				URL:       url,
			}

			createdImage, err := s.imageRepository.CreateImage(ctx, &image)
			if err != nil {
				return dto.ProductResponseDTO{}, fmt.Errorf("gagal menyimpan URL gambar ke db: %w", err)
			}

			createdProduct.ImageUrl = append(createdProduct.ImageUrl, createdImage)
		}
	}

	response := dto.ProductResponseDTO{
		ID:          createdProduct.ID,
		Name:        createdProduct.Name,
		Description: createdProduct.Description,
		Price:       createdProduct.Price,
		OldPrice:    createdProduct.OldPrice,
		Stock:       createdProduct.Stock,
		IsNew:       createdProduct.IsNew,
		Category:    createdProduct.Category,
		ImageUrl:    createdProduct.ImageUrl,
	}

	return response, nil
}

func (s *productService) UpdateProduct(ctx context.Context, productID uint64, productRequest dto.UpdateProductRequestDTO) (dto.ProductResponseDTO, error) {
	product := models.Product{
		ID:          productID,
		Name:        productRequest.Name,
		Description: productRequest.Description,
		Price:       productRequest.Price,
		OldPrice:    productRequest.OldPrice,
		Stock:       productRequest.Stock,
		IsNew:       productRequest.IsNew,
	}
	return s.productRepository.UpdateProduct(ctx, &product)
}

func (s *productService) DeleteProduct(ctx context.Context, productID uint64) error {
	return s.productRepository.DeleteProduct(ctx, productID)
}
