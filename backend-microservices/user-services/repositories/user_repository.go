package repositories

import (
	"context"

	"user-services/dto"
	"user-services/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAllUsers(ctx context.Context) ([]dto.UserResponseDTO, error)
	FindUserByUsername(ctx context.Context, username string) (dto.UserResponseDTO, error)
	CreateUser(ctx context.Context, user models.User) (dto.UserResponseDTO, error)
	UpdateUser(ctx context.Context, user models.User) (dto.UserResponseDTO, error)
	DeleteUser(ctx context.Context, user dto.UserDeleteDTO) error
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
		Select("id", "full_name", "avatar_url", "email", "member_since").
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) FindUserByUsername(ctx context.Context, username string) (dto.UserResponseDTO, error) {
	var user dto.UserResponseDTO
	if err := r.db.WithContext(ctx).
		Model(&models.User{}).
		Select("id", "full_name", "avatar_url", "email", "member_since").
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
		AvatarUrl:   user.AvatarUrl,
		Email:       user.Email,
		MemberSince: user.MemberSince,
	}, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, user dto.UserDeleteDTO) error {
	if err := r.db.WithContext(ctx).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
