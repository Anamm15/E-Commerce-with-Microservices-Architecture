package utils

import (
	"product-services/internal/dto"
	"product-services/internal/models"
	pb "product-services/pb"
)

func MapCategoryDTOResponseTogRPC(product dto.ProductResponseDTO) []*pb.CategoryResponse {
	var categories []*pb.CategoryResponse
	for _, category := range product.Category {
		categories = append(categories, &pb.CategoryResponse{
			Id:   category.ID,
			Name: category.Name,
		})
	}

	return categories
}

func MapImageDTOResponseTogRPC(product dto.ProductResponseDTO) []*pb.Image {
	var images []*pb.Image
	for _, image := range product.ImageUrl {
		images = append(images, &pb.Image{
			Id:  image.ID,
			Url: image.URL,
		})
	}

	return images
}

func MapCategoryModelsToDTO(product models.Product) []dto.CategoryResponseDTO {
	var categories []dto.CategoryResponseDTO
	for _, category := range product.Category {
		categories = append(categories, dto.CategoryResponseDTO{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	return categories
}

func MapImageModelsToDTO(product models.Product) []dto.ImageResponseDTO {
	var images []dto.ImageResponseDTO
	for _, image := range product.ImageUrl {
		images = append(images, dto.ImageResponseDTO{
			ID:  image.ID,
			URL: image.URL,
		})
	}

	return images
}
