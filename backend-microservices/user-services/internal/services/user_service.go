package services

import (
	"context"
	"errors"
	"fmt"

	"user-services/internal/constants"
	"user-services/internal/dto"
	"user-services/internal/repositories"
	"user-services/internal/utils"
)

type UserService interface {
	GetAllUsers(ctx context.Context) ([]dto.UserResponseDTO, error)
	GetUserByID(ctx context.Context, userId uint64) (dto.UserResponseDTO, error)
	GetUserByUsername(ctx context.Context, username string) (dto.UserResponseDTO, error)
	LoginUser(ctx context.Context, user dto.UserLoginDTO) (string, error)
	RegisterUser(ctx context.Context, user dto.UserCreateDTO) (dto.UserResponseDTO, error)
	UpdateUser(ctx context.Context, user dto.UserUpdateDTO) (dto.UserResponseDTO, error)
	DeleteUser(ctx context.Context, userId uint64) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) GetAllUsers(ctx context.Context) ([]dto.UserResponseDTO, error) {
	return s.userRepository.FindAllUsers(ctx)
}

func (s *userService) GetUserByID(ctx context.Context, userId uint64) (dto.UserResponseDTO, error) {
	user, err := s.userRepository.FindUserByID(ctx, userId)
	if err != nil {
		return dto.UserResponseDTO{}, err
	}
	return user, nil
}

func (s *userService) GetUserByUsername(ctx context.Context, username string) (dto.UserResponseDTO, error) {
	user, err := s.userRepository.FindUserByUsername(ctx, username)
	if err != nil {
		return dto.UserResponseDTO{}, err
	}
	return user, nil
}

func (s *userService) LoginUser(ctx context.Context, user dto.UserLoginDTO) (string, error) {
	userData, err := s.userRepository.FindUserEmail(ctx, user.Email)
	if err != nil {
		return "", errors.New(constants.ErrEmailPasswordIncorrect)
	}

	isMatch := utils.MatchPassword(userData.Password, user.Password)
	if !isMatch {
		return "", errors.New(constants.ErrEmailPasswordIncorrect)
	}

	fmt.Println(userData)
	token, err := utils.GenerateTokenJWT(userData.ID, userData.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) RegisterUser(ctx context.Context, user dto.UserCreateDTO) (dto.UserResponseDTO, error) {
	createdUser, err := s.userRepository.CreateUser(ctx, user.ToModel())
	if err != nil {
		return dto.UserResponseDTO{}, err
	}
	return createdUser, nil
}

func (s *userService) UpdateUser(ctx context.Context, user dto.UserUpdateDTO) (dto.UserResponseDTO, error) {
	updatedUser, err := s.userRepository.UpdateUser(ctx, user.ToModel())
	if err != nil {
		return dto.UserResponseDTO{}, err
	}
	return updatedUser, nil
}

func (s *userService) DeleteUser(ctx context.Context, userId uint64) error {
	return s.userRepository.DeleteUser(ctx, userId)
}
