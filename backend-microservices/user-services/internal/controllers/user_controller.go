package controllers

import (
	"context"

	"user-services/internal/dto"
	"user-services/internal/services"
	pb "user-services/pb"
)

type UserController struct {
	pb.UnimplementedUserServiceServer
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) CreateUser(ctx context.Context, req *pb.UserCreateRequest) (*pb.UserResponse, error) {
	userCreate := dto.UserCreateDTO{
		FullName: req.FullName,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	createdUser, err := c.service.CreateUser(ctx, userCreate)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id:        uint64(createdUser.ID),
		FullName:  createdUser.FullName,
		Username:  createdUser.Username,
		Email:     createdUser.Email,
		AvatarUrl: createdUser.AvatarUrl,
		Role:      createdUser.Role,
		// MemberSince: createdUser.MemberSince,
	}, nil
}

func (c *UserController) GetUserByUsername(ctx context.Context, req *pb.GetUserByUsernameRequest) (*pb.UserResponse, error) {
	user, err := c.service.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id:        uint64(user.ID),
		FullName:  user.FullName,
		Username:  user.Username,
		Email:     user.Email,
		AvatarUrl: user.AvatarUrl,
		Role:      user.Role,
		// MemberSince: timestamppb.New(user.CreatedAt),
	}, nil
}

func (c *UserController) GetUserByID(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	user, err := c.service.GetUserByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id:        uint64(user.ID),
		FullName:  user.FullName,
		Username:  user.Username,
		Email:     user.Email,
		AvatarUrl: user.AvatarUrl,
		Role:      user.Role,
		// MemberSince: timestamppb.New(user.CreatedAt),
	}, nil
}

func (c *UserController) GetAllUsers(ctx context.Context, _ *pb.Empty) (*pb.UserList, error) {
	users, err := c.service.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	var res pb.UserList
	for _, u := range users {
		res.Users = append(res.Users, &pb.UserResponse{
			Id:        uint64(u.ID),
			FullName:  u.FullName,
			Username:  u.Username,
			Email:     u.Email,
			AvatarUrl: u.AvatarUrl,
			Role:      u.Role,
			// MemberSince: timestamppb.New(u.CreatedAt),
		})
	}
	return &res, nil
}

func (c *UserController) UpdateUser(ctx context.Context, req *pb.UserUpdateRequest) (*pb.UserResponse, error) {
	userUpdate := dto.UserUpdateDTO{
		FullName:  req.FullName,
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
		AvatarUrl: req.AvatarUrl,
	}

	user, err := c.service.UpdateUser(ctx, userUpdate)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id:        uint64(user.ID),
		FullName:  user.FullName,
		Username:  user.Username,
		Email:     user.Email,
		AvatarUrl: user.AvatarUrl,
		Role:      user.Role,
		// MemberSince: timestamppb.New(user.CreatedAt),
	}, nil
}

func (c *UserController) DeleteUser(ctx context.Context, req *pb.UserDeleteRequest) (*pb.Empty, error) {
	err := c.service.DeleteUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
