package repositories

import (
	"context"

	"user-services/internal/dto"
	"user-services/internal/models"

	"gorm.io/gorm"
)

type AddressRepository interface {
	GetUserAddress(ctx context.Context, userId uint64) ([]dto.AddressResponseDTO, error)
	CreateUserAddress(ctx context.Context, address models.UserAddress) (dto.AddressResponseDTO, error)
	UpdateUserAddress(ctx context.Context, address models.UserAddress) (dto.AddressResponseDTO, error)
	DeleteUserAddress(ctx context.Context, addressId uint64, userId uint64) error
}

type addressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{
		db: db,
	}
}

func (r *addressRepository) GetUserAddress(ctx context.Context, userId uint64) ([]dto.AddressResponseDTO, error) {
	var address []dto.AddressResponseDTO

	if err := r.db.WithContext(ctx).
		Model(&models.UserAddress{}).
		Select("id", "user_id", "label", "recipient_name", "phone", "address", "city", "postal_code", "is_default").
		Where("user_id = ?", userId).
		Find(&address).Error; err != nil {
		return nil, err
	}

	return address, nil
}

func (r *addressRepository) CreateUserAddress(ctx context.Context, address models.UserAddress) (dto.AddressResponseDTO, error) {
	if err := r.db.WithContext(ctx).
		Model(&models.UserAddress{}).
		Select("id", "user_id", "label", "recipient_name", "phone", "address", "city", "postal_code", "is_default").
		Create(&address).Error; err != nil {
		return dto.AddressResponseDTO{}, err
	}

	return dto.AddressResponseDTO{
		ID:            address.ID,
		UserID:        address.UserID,
		Label:         address.Label,
		RecipientName: address.RecipientName,
		Phone:         address.Phone,
		Address:       address.Address,
		City:          address.City,
		PostalCode:    address.PostalCode,
		IsDefault:     address.IsDefault,
	}, nil
}

func (r *addressRepository) UpdateUserAddress(ctx context.Context, address models.UserAddress) (dto.AddressResponseDTO, error) {
	if err := r.db.WithContext(ctx).
		Model(&models.UserAddress{}).
		Select("id", "user_id", "label", "recipient_name", "phone", "address", "city", "postal_code", "is_default").
		Where("user_id = ?", address.UserID).
		Updates(&address).Error; err != nil {
		return dto.AddressResponseDTO{}, err
	}

	return dto.AddressResponseDTO{
		ID:            address.ID,
		UserID:        address.UserID,
		Label:         address.Label,
		RecipientName: address.RecipientName,
		Phone:         address.Phone,
		Address:       address.Address,
		City:          address.City,
		PostalCode:    address.PostalCode,
		IsDefault:     address.IsDefault,
	}, nil
}

func (r *addressRepository) DeleteUserAddress(ctx context.Context, addressId uint64, userId uint64) error {
	if err := r.db.WithContext(ctx).
		Model(&models.UserAddress{}).
		Where("user_id = ? AND id = ?", userId, addressId).
		Delete(&models.UserAddress{}).Error; err != nil {
		return err
	}

	return nil
}
