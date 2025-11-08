package repositories

import (
	"context"

	"user-services/internal/dto"
	"user-services/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAllUsers(ctx context.Context) ([]dto.UserResponseDTO, error)
	FindUserByUsername(ctx context.Context, username string) (dto.UserResponseDTO, error)
	CreateUser(ctx context.Context, user models.User) (dto.UserResponseDTO, error)
	UpdateUser(ctx context.Context, user models.User) (dto.UserResponseDTO, error)
	DeleteUser(ctx context.Context, userId uint64) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindAllUsers(ctx context.Context) ([]dto.UserResponseDTO, error) {
	var users []dto.UserResponseDTO

	if err := r.db.WithContext(ctx).
		Model(&models.User{}).
		Select("id", "full_name", "avatar_url", "email", "member_since", "username").
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) FindUserByUsername(ctx context.Context, username string) (dto.UserResponseDTO, error) {
	var user dto.UserResponseDTO
	if err := r.db.WithContext(ctx).
		Model(&models.User{}).
		Select("id", "full_name", "avatar_url", "email", "member_since", "username").
		Where("username = ?", username).
		First(&user).Error; err != nil {
		return dto.UserResponseDTO{}, err
	}

	return user, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user models.User) (dto.UserResponseDTO, error) {
	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return dto.UserResponseDTO{}, err
	}

	return dto.UserResponseDTO{
		ID:          user.ID,
		FullName:    user.FullName,
		AvatarUrl:   user.AvatarUrl,
		Username:    user.Username,
		Email:       user.Email,
		MemberSince: user.MemberSince,
	}, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user models.User) (dto.UserResponseDTO, error) {
	if err := r.db.WithContext(ctx).Save(&user).Error; err != nil {
		return dto.UserResponseDTO{}, err
	}
	return dto.UserResponseDTO{
		ID:          user.ID,
		FullName:    user.FullName,
		Username:    user.Username,
		AvatarUrl:   user.AvatarUrl,
		Email:       user.Email,
		MemberSince: user.MemberSince,
	}, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, userId uint64) error {
	if err := r.db.WithContext(ctx).Delete(&models.User{}, userId).Error; err != nil {
		return err
	}
	return nil
}
