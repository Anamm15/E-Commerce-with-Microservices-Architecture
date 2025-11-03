package controllers

import (
	"net/http"

	"user-services/dto"
	"user-services/services"
	"user-services/utils"

	"user-services/constants"

	"github.com/gin-gonic/gin"
)

type AddressController interface {
	GetUserAddress(ctx *gin.Context)
	CreateUserAddress(ctx *gin.Context)
	UpdateUserAddress(ctx *gin.Context)
	DeleteUserAddress(ctx *gin.Context)
}

type addressController struct {
	addressService services.AddressService
}

func NewAddressController(addressService services.AddressService) AddressController {
	return &addressController{
		addressService: addressService,
	}
}

func (c *addressController) GetUserAddress(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uint)

	address, err := c.addressService.GetUserAddress(ctx, userID)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ADDRESS_NOT_FOUND, err.Error(), ctx)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.ADDRESS_RETRIEVED_SUCCESSFULLY, address)
	ctx.JSON(http.StatusOK, res)
}

func (c *addressController) CreateUserAddress(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uint)

	var address dto.CreateAddressRequestDTO
	if err := ctx.ShouldBindJSON(&address); err != nil {
		res := utils.BuildResponseFailed(constants.INVALID_REQUEST, err.Error(), ctx)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	createdAddress, err := c.addressService.CreateUserAddress(ctx, userID, address)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ADDRESS_CREATION_FAILED, err.Error(), ctx)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.ADDRESS_CREATED_SUCCESSFULLY, createdAddress)
	ctx.JSON(http.StatusCreated, res)
}

func (c *addressController) UpdateUserAddress(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uint)

	var address dto.UpdateAddressRequestDTO
	if err := ctx.ShouldBindJSON(&address); err != nil {
		res := utils.BuildResponseFailed(constants.INVALID_REQUEST, err.Error(), ctx)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	updatedAddress, err := c.addressService.UpdateUserAddress(ctx, userID, address.ID, address)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ADDRESS_UPDATE_FAILED, err.Error(), ctx)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.ADDRESS_UPDATED_SUCCESSFULLY, updatedAddress)
	ctx.JSON(http.StatusOK, res)
}

func (c *addressController) DeleteUserAddress(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uint)
	addressIDParam := ctx.Param("address_id")

	addressID := utils.StringToUint(addressIDParam)
	if err := c.addressService.DeleteUserAddress(ctx, addressID, userID); err != nil {
		res := utils.BuildResponseFailed(constants.ADDRESS_DELETION_FAILED, err.Error(), ctx)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.ADDRESS_DELETED_SUCCESSFULLY, nil)
	ctx.JSON(http.StatusOK, res)
}
