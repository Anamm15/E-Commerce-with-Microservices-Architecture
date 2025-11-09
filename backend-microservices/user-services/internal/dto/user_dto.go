package dto

import (
	"time"

	"user-services/internal/models"
)

type UserDTO struct {
	ID        uint64    `json:"id"`
	FullName  string    `json:"full_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	AvatarUrl string    `json:"avatar_url"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponseDTO struct {
	ID          uint64    `json:"id"`
	FullName    string    `json:"full_name"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	AvatarUrl   string    `json:"avatar_url"`
	Role        string    `json:"role"`
	MemberSince time.Time `json:"member_since"`
}

type UserCreateDTO struct {
	FullName string `json:"full_name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserUpdateDTO struct {
	FullName  string `json:"full_name" binding:"omitempty"`
	Username  string `json:"username" binding:"omitempty"`
	Email     string `json:"email" binding:"omitempty,email"`
	Password  string `json:"password" binding:"omitempty,min=6"`
	AvatarUrl string `json:"avatar_url" binding:"omitempty,url"`
}

type UserDeleteDTO struct {
	Email string `json:"email" binding:"required,email"`
}

func (dto *UserCreateDTO) ToModel() models.User {
	return models.User{
		FullName: dto.FullName,
		Username: dto.Username,
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func (dto *UserUpdateDTO) ToModel() models.User {
	return models.User{
		FullName:  dto.FullName,
		Email:     dto.Email,
		Username:  dto.Username,
		Password:  dto.Password,
		AvatarUrl: dto.AvatarUrl,
	}
}
