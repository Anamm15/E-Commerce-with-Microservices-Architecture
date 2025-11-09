package controllers

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"user-services/internal/constants"
	"user-services/internal/dto"
	"user-services/internal/services"
	pb "user-services/pb"
)

type AddressController struct {
	pb.UnimplementedAddressServiceServer
	addressService services.AddressService
}

func NewAddressController(addressService services.AddressService) *AddressController {
	return &AddressController{
		addressService: addressService,
	}
}

func (c *AddressController) GetUserAddress(ctx context.Context, req *pb.GetAddressRequest) (*pb.AddressListResponse, error) {
	res, err := c.addressService.GetUserAddress(context.Background(), req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, constants.ErrAddressServiceGet, err)
	}

	var pbRes *pb.AddressListResponse
	for _, address := range res {
		pbRes.Addresses = append(pbRes.Addresses, &pb.AddressResponse{
			Id:            address.ID,
			Label:         address.Label,
			Address:       address.Address,
			City:          address.City,
			IsDefault:     address.IsDefault,
			PostalCode:    address.PostalCode,
			Phone:         address.Phone,
			RecipientName: address.RecipientName,
		})
	}

	return pbRes, nil
}

func (c *AddressController) CreateUserAddress(ctx context.Context, req *pb.CreateAddressRequest) (*pb.AddressResponse, error) {
	reqBody := dto.CreateAddressRequestDTO{
		Label:         req.Label,
		Address:       req.Address,
		City:          req.City,
		PostalCode:    req.PostalCode,
		Phone:         req.Phone,
		RecipientName: req.RecipientName,
	}

	createdAddress, err := c.addressService.CreateUserAddress(context.Background(), req.GetUserId(), reqBody)
	if err != nil {
		return nil, status.Errorf(codes.Internal, constants.ErrAddressServiceCreate, err)
	}

	return &pb.AddressResponse{
		Id:            createdAddress.ID,
		Label:         createdAddress.Label,
		Address:       createdAddress.Address,
		City:          createdAddress.City,
		IsDefault:     createdAddress.IsDefault,
		PostalCode:    createdAddress.PostalCode,
		Phone:         createdAddress.Phone,
		RecipientName: createdAddress.RecipientName,
	}, nil
}

func (c *AddressController) UpdateUserAddress(ctx context.Context, req *pb.UpdateAddressRequest) (*pb.AddressResponse, error) {
	reqBody := dto.UpdateAddressRequestDTO{
		Label:         req.Label,
		Address:       req.Address,
		City:          req.City,
		IsDefault:     req.IsDefault,
		PostalCode:    req.PostalCode,
		Phone:         req.Phone,
		RecipientName: req.RecipientName,
	}

	updatedAddress, err := c.addressService.UpdateUserAddress(ctx, req.Id, req.UserId, reqBody)
	if err != nil {
		return nil, status.Errorf(codes.Internal, constants.ErrAddressServiceUpdate, err)
	}

	return &pb.AddressResponse{
		Id:            updatedAddress.ID,
		Label:         updatedAddress.Label,
		Address:       updatedAddress.Address,
		City:          updatedAddress.City,
		IsDefault:     updatedAddress.IsDefault,
		PostalCode:    updatedAddress.PostalCode,
		Phone:         updatedAddress.Phone,
		RecipientName: updatedAddress.RecipientName,
	}, nil
}

func (c *AddressController) DeleteUserAddress(ctx context.Context, req *pb.AddressIDRequest) (*pb.Empty, error) {
	err := c.addressService.DeleteUserAddress(ctx, req.AddressId, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, constants.ErrAddressServiceDelete, err)
	}

	return &pb.Empty{}, nil
}
