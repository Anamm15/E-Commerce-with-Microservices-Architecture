package repositories

import (
	"context"
	"errors"

	"product-services/internal/dto"
	"product-services/internal/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAllCategories(ctx context.Context) ([]dto.CategoryResponseDTO, error)
	AddProductCategory(ctx context.Context, productID uint, categoryIDs []uint32) ([]dto.CategoryResponseDTO, error)
	CreateCategory(ctx context.Context, category *models.Category) (dto.CategoryResponseDTO, error)
	UpdateCategory(ctx context.Context, category *models.Category) (dto.CategoryResponseDTO, error)
	DeleteCategory(ctx context.Context, categoryId uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) GetAllCategories(ctx context.Context) ([]dto.CategoryResponseDTO, error) {
	var categories []dto.CategoryResponseDTO
	if err := r.db.WithContext(ctx).
		Model(&models.Category{}).
		Select("id", "name").
		Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) AddProductCategory(ctx context.Context, productID uint, categoryIDs []uint32) ([]dto.CategoryResponseDTO, error) {
	var categories []models.Category

	if err := r.db.WithContext(ctx).
		Where("id IN ?", categoryIDs).
		Find(&categories).Error; err != nil {
		return nil, err
	}

	if len(categories) == 0 {
		return nil, errors.New("kategori tidak ditemukan")
	}

	var product models.Product
	product.ID = productID

	if err := r.db.WithContext(ctx).
		Model(&product).
		Association("Category").
		Append(&categories); err != nil {
		return nil, err
	}

	var categoryResponseDTOs []dto.CategoryResponseDTO
	for _, category := range categories {
		categoryResponseDTOs = append(categoryResponseDTOs, dto.CategoryResponseDTO{ID: category.ID, Name: category.Name})
	}
	return categoryResponseDTOs, nil
}

func (r *categoryRepository) CreateCategory(ctx context.Context, category *models.Category) (dto.CategoryResponseDTO, error) {
	if err := r.db.Create(&category).Error; err != nil {
		return dto.CategoryResponseDTO{}, err
	}
	return dto.CategoryResponseDTO{ID: category.ID, Name: category.Name}, nil
}

func (r *categoryRepository) UpdateCategory(ctx context.Context, category *models.Category) (dto.CategoryResponseDTO, error) {
	if err := r.db.Save(&category).Error; err != nil {
		return dto.CategoryResponseDTO{}, err
	}
	return dto.CategoryResponseDTO{ID: category.ID, Name: category.Name}, nil
}

func (r *categoryRepository) DeleteCategory(ctx context.Context, categoryId uint) error {
	if err := r.db.Where("id = ?", categoryId).Delete(&models.Category{}).Error; err != nil {
		return err
	}
	return nil
}
