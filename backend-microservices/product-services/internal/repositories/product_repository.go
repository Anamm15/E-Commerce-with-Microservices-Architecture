package repositories

import (
	"context"

	"product-services/internal/dto"
	"product-services/internal/models"
	"product-services/internal/utils"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProducts(ctx context.Context) ([]dto.ProductResponseDTO, error)
	GetProductById(ctx context.Context, productId uint64) (dto.ProductResponseDTO, error)
	GetProductsByCategory(ctx context.Context, categoryId uint64) ([]dto.ProductResponseDTO, error)
	CreateProduct(ctx context.Context, product *models.Product) (dto.ProductResponseDTO, error)
	UpdateProduct(ctx context.Context, product *models.Product) (dto.ProductResponseDTO, error)
	DeleteProduct(ctx context.Context, productId uint64) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetAllProducts(ctx context.Context) ([]dto.ProductResponseDTO, error) {
	var products []models.Product

	if err := r.db.WithContext(ctx).
		Preload("Category").
		Preload("ImageUrl").
		Find(&products).Error; err != nil {
		return nil, err
	}

	var productDTOs []dto.ProductResponseDTO
	for _, product := range products {
		categoryDTOs := utils.MapCategoryModelsToDTO(product)
		imageDTOs := utils.MapImageModelsToDTO(product)

		productDTOs = append(productDTOs, dto.ProductResponseDTO{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			OldPrice:    product.OldPrice,
			Stock:       product.Stock,
			Rating:      product.Rating,
			IsNew:       product.IsNew,
			Category:    categoryDTOs,
			ImageUrl:    imageDTOs,
		})
	}

	return productDTOs, nil
}

func (r *productRepository) GetProductById(ctx context.Context, productId uint64) (dto.ProductResponseDTO, error) {
	var product models.Product
	if err := r.db.WithContext(ctx).
		Preload("Category").
		Preload("ImageUrl").
		Where("id = ?", productId).
		Find(&product).Error; err != nil {
		return dto.ProductResponseDTO{}, err
	}

	categoryDTOs := utils.MapCategoryModelsToDTO(product)
	imageDTOs := utils.MapImageModelsToDTO(product)
	return dto.ProductResponseDTO{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		OldPrice:    product.OldPrice,
		Stock:       product.Stock,
		Rating:      product.Rating,
		IsNew:       product.IsNew,
		Category:    categoryDTOs,
		ImageUrl:    imageDTOs,
	}, nil
}

func (r *productRepository) GetProductsByCategory(ctx context.Context, categoryId uint64) ([]dto.ProductResponseDTO, error) {
	var products []models.Product
	if err := r.db.WithContext(ctx).
		Preload("Category").
		Preload("ImageUrl").
		Where("category_id = ?", categoryId).
		Find(&products).Error; err != nil {
		return nil, err
	}

	var productDTOs []dto.ProductResponseDTO
	for _, product := range products {
		categoryDTOs := utils.MapCategoryModelsToDTO(product)
		imageDTOs := utils.MapImageModelsToDTO(product)

		productDTOs = append(productDTOs, dto.ProductResponseDTO{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			OldPrice:    product.OldPrice,
			Stock:       product.Stock,
			Rating:      product.Rating,
			IsNew:       product.IsNew,
			Category:    categoryDTOs,
			ImageUrl:    imageDTOs,
		})
	}

	return productDTOs, nil
}

func (r *productRepository) CreateProduct(ctx context.Context, product *models.Product) (dto.ProductResponseDTO, error) {
	if err := r.db.WithContext(ctx).
		Create(&product).Error; err != nil {
		return dto.ProductResponseDTO{}, err
	}

	return dto.ProductResponseDTO{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		OldPrice:    product.OldPrice,
		Stock:       product.Stock,
		Rating:      product.Rating,
		IsNew:       product.IsNew,
	}, nil
}

func (r *productRepository) UpdateProduct(ctx context.Context, product *models.Product) (dto.ProductResponseDTO, error) {
	if err := r.db.WithContext(ctx).
		Save(&product).Error; err != nil {
		return dto.ProductResponseDTO{}, err
	}

	return dto.ProductResponseDTO{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		OldPrice:    product.OldPrice,
		Stock:       product.Stock,
		Rating:      product.Rating,
		IsNew:       product.IsNew,
	}, nil
}

func (r *productRepository) DeleteProduct(ctx context.Context, productId uint64) error {
	if err := r.db.WithContext(ctx).
		Where("id = ?", productId).
		Delete(&models.Product{}).Error; err != nil {
		return err
	}
	return nil
}
