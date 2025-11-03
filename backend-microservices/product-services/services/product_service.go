package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"product-services/dto"
	"product-services/models"
	"product-services/repositories"
	"product-services/storages"
)

type ProductService interface {
	GetAllProducts(ctx context.Context) ([]dto.ProductResponseDTO, error)
	GetProductByID(ctx context.Context, productID uint) (dto.ProductResponseDTO, error)
	GetProductByCategoryID(ctx context.Context, categoryID uint) ([]dto.ProductResponseDTO, error)
	CreateProduct(ctx context.Context, productRequest dto.CreateProductRequestDTO) (dto.ProductResponseDTO, error)
	UpdateProduct(ctx context.Context, productID uint, productRequest dto.UpdateProductRequestDTO) (dto.ProductResponseDTO, error)
	DeleteProduct(ctx context.Context, productID uint) error
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

func (s *productService) GetProductByID(ctx context.Context, productID uint) (dto.ProductResponseDTO, error) {
	return s.productRepository.GetProductById(ctx, productID)
}

func (s *productService) GetProductByCategoryID(ctx context.Context, categoryID uint) ([]dto.ProductResponseDTO, error) {
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
		for _, fileHeader := range productRequest.Images {
			file, err := fileHeader.Open()
			if err != nil {
				return dto.ProductResponseDTO{}, fmt.Errorf("gagal membuka file %s: %w", fileHeader.Filename, err)
			}

			uploadErr := func() error {
				defer file.Close()

				path := fmt.Sprintf("products/%d/%d_%s", createdProduct.ID, time.Now().UnixNano(), fileHeader.Filename)
				contentType := fileHeader.Header.Get("Content-Type")

				if contentType == "" {
					buf := make([]byte, 512)
					n, err := file.Read(buf)
					if err != nil && err != io.EOF {
						return fmt.Errorf("gagal membaca file untuk deteksi content-type: %w", err)
					}

					contentType = http.DetectContentType(buf[:n])

					if _, err = file.Seek(0, io.SeekStart); err != nil {
						return fmt.Errorf("gagal reset file reader: %w", err)
					}
				}

				url, err := s.storage.UploadFile(ctx, path, file, contentType)
				if err != nil {
					return fmt.Errorf("gagal upload file: %w", err)
				}

				image := models.Image{
					ProductID: createdProduct.ID,
					URL:       url,
				}

				var createdImage dto.ImageResponseDTO
				createdImage, err = s.imageRepository.CreateImage(ctx, &image)
				if err != nil {
					return fmt.Errorf("gagal menyimpan URL gambar ke db: %w", err)
				}

				createdProduct.ImageUrl = append(createdProduct.ImageUrl, createdImage)
				return nil
			}()

			if uploadErr != nil {
				log.Printf("Gagal meng-upload file: %v", uploadErr)
				return dto.ProductResponseDTO{}, uploadErr
			}
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

func (s *productService) UpdateProduct(ctx context.Context, productID uint, productRequest dto.UpdateProductRequestDTO) (dto.ProductResponseDTO, error) {
	product := models.Product{Name: productRequest.Name}
	return s.productRepository.UpdateProduct(ctx, &product)
}

func (s *productService) DeleteProduct(ctx context.Context, productID uint) error {
	return s.productRepository.DeleteProduct(ctx, productID)
}
