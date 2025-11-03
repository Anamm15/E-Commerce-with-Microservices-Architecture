package services

import (
	"context"

	"user-services/dto"
	"user-services/models"
	"user-services/repositories"
)

type AddressService interface {
	GetUserAddress(ctx context.Context, userID uint) ([]dto.AddressResponseDTO, error)
	CreateUserAddress(ctx context.Context, userId uint, address dto.CreateAddressRequestDTO) (dto.AddressResponseDTO, error)
	UpdateUserAddress(ctx context.Context, userId uint, addressId uint, address dto.UpdateAddressRequestDTO) (dto.AddressResponseDTO, error)
	DeleteUserAddress(ctx context.Context, addressId uint, userId uint) error
}

type addressService struct {
	addressRepository repositories.AddressRepository
}

func NewAddressService(addressRepository repositories.AddressRepository) AddressService {
	return &addressService{
		addressRepository: addressRepository,
	}
}

func (s *addressService) GetUserAddress(ctx context.Context, userID uint) ([]dto.AddressResponseDTO, error) {
	return s.addressRepository.GetUserAddress(ctx, userID)
}

func (s *addressService) CreateUserAddress(ctx context.Context, userId uint, address dto.CreateAddressRequestDTO) (dto.AddressResponseDTO, error) {
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

func (s *addressService) UpdateUserAddress(ctx context.Context, userId uint, addressId uint, address dto.UpdateAddressRequestDTO) (dto.AddressResponseDTO, error) {
	addressInput := models.UserAddress{
		UserID:        userId,
		Label:         address.Label,
		RecipientName: address.RecipientName,
		Phone:         address.Phone,
		Address:       address.Address,
		City:          address.City,
		PostalCode:    address.PostalCode,
	}

	updatedAddress, err := s.addressRepository.UpdateUserAddress(ctx, addressId, addressInput)
	if err != nil {
		return dto.AddressResponseDTO{}, err
	}

	return updatedAddress, nil
}

func (s *addressService) DeleteUserAddress(ctx context.Context, addressId uint, userId uint) error {
	return s.addressRepository.DeleteUserAddress(ctx, addressId, userId)
}
