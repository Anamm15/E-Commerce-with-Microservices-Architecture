package services

import (
	"context"

	"user-services/internal/dto"
	"user-services/internal/models"
	"user-services/internal/repositories"
)

type AddressService interface {
	GetUserAddress(ctx context.Context, userID uint64) ([]dto.AddressResponseDTO, error)
	CreateUserAddress(ctx context.Context, userId uint64, address dto.CreateAddressRequestDTO) (dto.AddressResponseDTO, error)
	UpdateUserAddress(ctx context.Context, userId uint64, addressId uint64, address dto.UpdateAddressRequestDTO) (dto.AddressResponseDTO, error)
	DeleteUserAddress(ctx context.Context, addressId uint64, userId uint64) error
}

type addressService struct {
	addressRepository repositories.AddressRepository
}

func NewAddressService(addressRepository repositories.AddressRepository) AddressService {
	return &addressService{
		addressRepository: addressRepository,
	}
}

func (s *addressService) GetUserAddress(ctx context.Context, userID uint64) ([]dto.AddressResponseDTO, error) {
	return s.addressRepository.GetUserAddress(ctx, userID)
}

func (s *addressService) CreateUserAddress(ctx context.Context, userId uint64, address dto.CreateAddressRequestDTO) (dto.AddressResponseDTO, error) {
	addressInput := models.UserAddress{
		UserID:        userId,
		Label:         address.Label,
		RecipientName: address.RecipientName,
		Phone:         address.Phone,
		Address:       address.Address,
		City:          address.City,
		PostalCode:    address.PostalCode,
	}

	createdAddress, err := s.addressRepository.CreateUserAddress(ctx, addressInput)
	if err != nil {
		return dto.AddressResponseDTO{}, err
	}

	return createdAddress, nil
}

func (s *addressService) UpdateUserAddress(ctx context.Context, userId uint64, addressId uint64, address dto.UpdateAddressRequestDTO) (dto.AddressResponseDTO, error) {
	addressInput := models.UserAddress{
		ID:            addressId,
		UserID:        userId,
		Label:         address.Label,
		RecipientName: address.RecipientName,
		Phone:         address.Phone,
		Address:       address.Address,
		City:          address.City,
		PostalCode:    address.PostalCode,
	}

	updatedAddress, err := s.addressRepository.UpdateUserAddress(ctx, addressInput)
	if err != nil {
		return dto.AddressResponseDTO{}, err
	}

	return updatedAddress, nil
}

func (s *addressService) DeleteUserAddress(ctx context.Context, addressId uint64, userId uint64) error {
	return s.addressRepository.DeleteUserAddress(ctx, addressId, userId)
}
