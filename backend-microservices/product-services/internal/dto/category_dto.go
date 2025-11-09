package dto

type CategoryResponseDTO struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type CreateCategoryRequestDTO struct {
	Name string `json:"name" validate:"required"`
}

type UpdateCategoryRequestDTO struct {
	Name string `json:"name" validate:"required"`
}
