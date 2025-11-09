package user

import (
	"time"
)

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
