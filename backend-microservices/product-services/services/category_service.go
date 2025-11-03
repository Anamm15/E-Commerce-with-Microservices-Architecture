package services

import (
	"context"

	"product-services/dto"
	"product-services/models"
	"product-services/repositories"
)

type CategoryService interface {
	GetAllCategories(ctx context.Context) ([]dto.CategoryResponseDTO, error)
	CreateCategory(ctx context.Context, categoryRequest dto.CreateCategoryRequestDTO) (dto.CategoryResponseDTO, error)
	UpdateCategory(ctx context.Context, categoryID uint, categoryRequest dto.UpdateCategoryRequestDTO) (dto.CategoryResponseDTO, error)
	DeleteCategory(ctx context.Context, categoryID uint) error
}

type categoryService struct {
	categoryRepository repositories.CategoryRepository
}

func NewCategoryService(categoryRepository repositories.CategoryRepository) CategoryService {
	return &categoryService{categoryRepository: categoryRepository}
}

func (s *categoryService) GetAllCategories(ctx context.Context) ([]dto.CategoryResponseDTO, error) {
	categories, err := s.categoryRepository.GetAllCategories(ctx)
	if err != nil {
		return nil, err
	}
	var categoryResponseDTOs []dto.CategoryResponseDTO
	for _, category := range categories {
		categoryResponseDTOs = append(categoryResponseDTOs, dto.CategoryResponseDTO{ID: category.ID, Name: category.Name})
	}
	return categoryResponseDTOs, nil
}

func (s *categoryService) CreateCategory(ctx context.Context, categoryRequest dto.CreateCategoryRequestDTO) (dto.CategoryResponseDTO, error) {
	category := models.Category{Name: categoryRequest.Name}
	categoryResponseDTO, err := s.categoryRepository.CreateCategory(ctx, &category)
	if err != nil {
		return dto.CategoryResponseDTO{}, err
	}
	return categoryResponseDTO, nil
}

func (s *categoryService) UpdateCategory(ctx context.Context, categoryID uint, categoryRequest dto.UpdateCategoryRequestDTO) (dto.CategoryResponseDTO, error) {
	category := models.Category{Name: categoryRequest.Name}
	categoryResponseDTO, err := s.categoryRepository.UpdateCategory(ctx, &category)
	if err != nil {
		return dto.CategoryResponseDTO{}, err
	}
	return categoryResponseDTO, nil
}

func (s *categoryService) DeleteCategory(ctx context.Context, categoryID uint) error {
	return s.categoryRepository.DeleteCategory(ctx, categoryID)
}
