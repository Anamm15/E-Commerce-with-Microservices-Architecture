package controllers

import (
	"context"
	"net/http"
	"strconv"

	"user-services/internal/constants"
	"user-services/internal/dto"
	"user-services/internal/services"
	"user-services/internal/utils"
	pb "user-services/pb"

	"github.com/gin-gonic/gin"
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

func (c *AddressController) GetUserAddress(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uint64)

	req := &pb.GetAddressRequest{
		UserId: userID,
	}

	res, err := c.addressService.GetUserAddress(context.Background(), uint(req.UserId))
	if err != nil {
		resp := utils.BuildResponseFailed(constants.ADDRESS_NOT_FOUND, err.Error(), ctx)
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := utils.BuildResponseSuccess(constants.ADDRESS_RETRIEVED_SUCCESSFULLY, res)
	ctx.JSON(http.StatusOK, resp)
}

func (c *AddressController) CreateUserAddress(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uint64)

	var reqBody pb.CreateAddressRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		resp := utils.BuildResponseFailed(constants.INVALID_REQUEST, err.Error(), ctx)
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	req := dto.CreateAddressRequestDTO{
		Label:         reqBody.Label,
		Address:       reqBody.Address,
		City:          reqBody.City,
		PostalCode:    reqBody.PostalCode,
		Phone:         reqBody.Phone,
		RecipientName: reqBody.RecipientName,
	}

	res, err := c.addressService.CreateUserAddress(context.Background(), uint(userID), req)
	if err != nil {
		resp := utils.BuildResponseFailed(constants.ADDRESS_CREATION_FAILED, err.Error(), ctx)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := utils.BuildResponseSuccess(constants.ADDRESS_CREATED_SUCCESSFULLY, res)
	ctx.JSON(http.StatusCreated, resp)
}

func (c *AddressController) UpdateUserAddress(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uint)

	var reqBody pb.UpdateAddressRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		resp := utils.BuildResponseFailed(constants.INVALID_REQUEST, err.Error(), ctx)
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	req := dto.UpdateAddressRequestDTO{
		Label:         reqBody.Label,
		Address:       reqBody.Address,
		City:          reqBody.City,
		IsDefault:     reqBody.IsDefault,
		PostalCode:    reqBody.PostalCode,
		Phone:         reqBody.Phone,
		RecipientName: reqBody.RecipientName,
	}

	updatedAddress, err := c.addressService.UpdateUserAddress(context.Background(), uint(reqBody.Id), userID, req)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ADDRESS_UPDATE_FAILED, err.Error(), ctx)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.ADDRESS_UPDATED_SUCCESSFULLY, updatedAddress)
	ctx.JSON(http.StatusOK, res)
}

func (c *AddressController) DeleteUserAddress(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uint64)
	addressIDParam := ctx.Param("address_id")

	addressID, err := strconv.ParseUint(addressIDParam, 10, 64)
	if err != nil {
		resp := utils.BuildResponseFailed(constants.INVALID_REQUEST, "invalid address id", ctx)
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	req := &pb.AddressIDRequest{
		AddressId: addressID,
		UserId:    userID,
	}

	err = c.addressService.DeleteUserAddress(context.Background(), uint(req.AddressId), uint(req.UserId))
	if err != nil {
		res := utils.BuildResponseFailed(constants.ADDRESS_DELETION_FAILED, err.Error(), ctx)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.ADDRESS_DELETED_SUCCESSFULLY, nil)
	ctx.JSON(http.StatusOK, res)
}
