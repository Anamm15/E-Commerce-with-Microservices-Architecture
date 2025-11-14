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
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdateDTO struct {
	FullName  string `json:"full_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	AvatarUrl string `json:"avatar_url"`
}

type UserDeleteDTO struct {
	Email string `json:"email" binding:"required,email"`
}

type UserLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserEmailResponseDTO struct {
	ID       uint64 `json:"id"`
	Role     string `json:"role"`
	Password string `json:"password"`
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
