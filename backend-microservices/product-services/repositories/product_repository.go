package repositories

import (
	"context"

	"product-services/dto"
	"product-services/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProducts(ctx context.Context) ([]dto.ProductResponseDTO, error)
	GetProductById(ctx context.Context, productId uint) (dto.ProductResponseDTO, error)
	GetProductsByCategory(ctx context.Context, categoryId uint) ([]dto.ProductResponseDTO, error)
	CreateProduct(ctx context.Context, product *models.Product) (dto.ProductResponseDTO, error)
	UpdateProduct(ctx context.Context, product *models.Product) (dto.ProductResponseDTO, error)
	DeleteProduct(ctx context.Context, productId uint) error
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
	for _, p := range products {
		var imageDTOs []dto.ImageResponseDTO
		for _, img := range p.ImageUrl {
			imageDTOs = append(imageDTOs, dto.ImageResponseDTO{
				ID:  img.ID,
				URL: img.URL,
			})
		}

		var categoryDTOs []dto.CategoryResponseDTO
		for _, c := range p.Category {
			categoryDTOs = append(categoryDTOs, dto.CategoryResponseDTO{
				ID:   c.ID,
				Name: c.Name,
			})
		}

		productDTOs = append(productDTOs, dto.ProductResponseDTO{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			OldPrice:    p.OldPrice,
			Stock:       p.Stock,
			Rating:      p.Rating,
			IsNew:       p.IsNew,
			Category:    categoryDTOs,
			ImageUrl:    imageDTOs,
		})
	}

	return productDTOs, nil
}

func (r *productRepository) GetProductById(ctx context.Context, productId uint) (dto.ProductResponseDTO, error) {
	var product models.Product
	if err := r.db.WithContext(ctx).
		Preload("Category").
		Preload("ImageUrl").
		Where("id = ?", productId).
		Find(&product).Error; err != nil {
		return dto.ProductResponseDTO{}, err
	}

	var imageDTOs []dto.ImageResponseDTO
	for _, img := range product.ImageUrl {
		imageDTOs = append(imageDTOs, dto.ImageResponseDTO{
			ID:  img.ID,
			URL: img.URL,
		})
	}

	var categoryDTOs []dto.CategoryResponseDTO
	for _, c := range product.Category {
		categoryDTOs = append(categoryDTOs, dto.CategoryResponseDTO{
			ID:   c.ID,
			Name: c.Name,
		})
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
		Category:    categoryDTOs,
		ImageUrl:    imageDTOs,
	}, nil
}

func (r *productRepository) GetProductsByCategory(ctx context.Context, categoryId uint) ([]dto.ProductResponseDTO, error) {
	var products []models.Product
	if err := r.db.WithContext(ctx).
		Preload("Category").
		Preload("ImageUrl").
		Where("category_id = ?", categoryId).
		Find(&products).Error; err != nil {
		return nil, err
	}

	var productDTOs []dto.ProductResponseDTO
	for _, p := range products {
		var imageDTOs []dto.ImageResponseDTO
		for _, img := range p.ImageUrl {
			imageDTOs = append(imageDTOs, dto.ImageResponseDTO{
				ID:  img.ID,
				URL: img.URL,
			})
		}

		var categoryDTOs []dto.CategoryResponseDTO
		for _, c := range p.Category {
			categoryDTOs = append(categoryDTOs, dto.CategoryResponseDTO{
				ID:   c.ID,
				Name: c.Name,
			})
		}

		productDTOs = append(productDTOs, dto.ProductResponseDTO{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			OldPrice:    p.OldPrice,
			Stock:       p.Stock,
			Rating:      p.Rating,
			IsNew:       p.IsNew,
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
	if err := r.db.Save(&product).Error; err != nil {
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

func (r *productRepository) DeleteProduct(ctx context.Context, productId uint) error {
	if err := r.db.Where("id = ?", productId).Delete(&models.Product{}).Error; err != nil {
		return err
	}
	return nil
}
