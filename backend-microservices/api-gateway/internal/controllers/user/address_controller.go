package user

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"api-gateway/internal/constants"
	dto "api-gateway/internal/dto/user"
	userpb "api-gateway/internal/pb/user"
	"api-gateway/internal/utils"

	"github.com/gin-gonic/gin"
)

type AddressController struct {
	UserClient userpb.AddressServiceClient
}

func NewAddressController(UserClient userpb.AddressServiceClient) *AddressController {
	return &AddressController{
		UserClient: UserClient,
	}
}

func (c *AddressController) GetUserAddress(ctx *gin.Context) {
	userID := ctx.MustGet(constants.ContextKeyUserID).(uint)

	req := &userpb.GetAddressRequest{
		UserId: uint64(userID),
	}

	userAddresses, err := c.UserClient.GetUserAddress(context.Background(), req)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ErrGetAddress, err.Error(), userAddresses)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.SuccessAddressRetrieved, userAddresses)
	ctx.JSON(http.StatusOK, res)
}

func (c *AddressController) CreateUserAddress(ctx *gin.Context) {
	userID := ctx.MustGet(constants.ContextKeyUserID).(uint)

	var reqBody dto.CreateAddressRequestDTO

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		res := utils.BuildResponseFailed(constants.ErrInvalidRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	req := &userpb.CreateAddressRequest{
		UserId:        uint64(userID),
		Label:         reqBody.Label,
		RecipientName: reqBody.RecipientName,
		Phone:         reqBody.Phone,
		Address:       reqBody.Address,
		City:          reqBody.City,
		PostalCode:    reqBody.PostalCode,
	}

	ctxGrpc, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	createdAddress, err := c.UserClient.CreateUserAddress(ctxGrpc, req)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ErrCreateAddresss, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.SuccessAddressCreated, createdAddress)
	ctx.JSON(http.StatusCreated, res)
}

func (c *AddressController) UpdateUserAddress(ctx *gin.Context) {
	userID := ctx.MustGet(constants.ContextKeyUserID).(uint64)
	addressIDParam := ctx.Param(constants.ParamAddressID)
	addressID := utils.StringToUint(addressIDParam)

	var reqBody dto.UpdateAddressRequestDTO
	reqBody.ID = addressID
	reqBody.UserID = userID

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		res := utils.BuildResponseFailed(constants.ErrInvalidRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	req := &userpb.UpdateAddressRequest{
		Id:            uint64(reqBody.ID),
		UserId:        uint64(userID),
		Label:         reqBody.Label,
		RecipientName: reqBody.RecipientName,
		Phone:         reqBody.Phone,
		Address:       reqBody.Address,
		City:          reqBody.City,
		PostalCode:    reqBody.PostalCode,
		IsDefault:     reqBody.IsDefault,
	}

	ctxGrpc, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updatedAddress, err := c.UserClient.UpdateUserAddress(ctxGrpc, req)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ErrUpdateAddress, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.SuccessAddressUpdated, updatedAddress)
	ctx.JSON(http.StatusOK, res)
}

func (c *AddressController) DeleteUserAddress(ctx *gin.Context) {
	userID := ctx.MustGet(constants.ContextKeyUserID).(uint)
	addressIDParam := ctx.Param(constants.ParamAddressID)

	addressID, _ := strconv.Atoi(addressIDParam)

	req := &userpb.AddressIDRequest{
		UserId:    uint64(userID),
		AddressId: uint64(addressID),
	}

	ctxGrpc, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := c.UserClient.DeleteUserAddress(ctxGrpc, req)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ErrDeleteAddress, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.SuccessAddressDeleted, nil)
	ctx.JSON(http.StatusOK, res)
}
