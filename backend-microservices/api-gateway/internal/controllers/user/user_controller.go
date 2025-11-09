package user

import (
	"context"
	"net/http"

	"api-gateway/internal/constants"
	dto "api-gateway/internal/dto/user"
	userpb "api-gateway/internal/pb/user"
	"api-gateway/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserClient userpb.UserServiceClient
}

func NewUserController(client userpb.UserServiceClient) *UserController {
	return &UserController{UserClient: client}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var req dto.UserCreateDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildResponseFailed(constants.ErrInvalidRequest, err.Error(), nil))
		return
	}

	grpcReq := &userpb.UserCreateRequest{
		FullName: req.FullName,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	resp, err := uc.UserClient.CreateUser(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildResponseFailed(constants.ErrCreateUser, err.Error(), nil))
		return
	}

	c.JSON(http.StatusCreated, utils.BuildResponseSuccess(constants.SuccessUserCreated, resp))
}

func (uc *UserController) GetUserByUsername(c *gin.Context) {
	username := c.Query(constants.QueryUsername)

	grpcReq := &userpb.GetUserByUsernameRequest{Username: username}

	resp, err := uc.UserClient.GetUserByUsername(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildResponseFailed(constants.ErrGetUser, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, utils.BuildResponseSuccess(constants.SuccessUserFetched, resp))
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	resp, err := uc.UserClient.GetAllUsers(context.Background(), &userpb.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildResponseFailed(constants.ErrGetUsers, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, utils.BuildResponseSuccess(constants.SuccessUsersFetched, resp.Users))
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	idParam := c.Param(constants.ParamID)
	if idParam == "" {
		c.JSON(http.StatusBadRequest, utils.BuildResponseFailed(constants.ErrInvalidRequest, constants.ErrIDRequired, nil))
		return
	}

	reqID := utils.StringToUint(idParam)
	if reqID == 0 {
		c.JSON(http.StatusBadRequest, utils.BuildResponseFailed(constants.ErrInvalidRequest, constants.ErrInvalidID, nil))
		return
	}

	var req dto.UserUpdateDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildResponseFailed(constants.ErrInvalidRequest, err.Error(), nil))
		return
	}

	grpcReq := &userpb.UserUpdateRequest{
		Id:        uint64(reqID),
		FullName:  req.FullName,
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
		AvatarUrl: req.AvatarUrl,
	}

	resp, err := uc.UserClient.UpdateUser(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildResponseFailed(constants.ErrUpdateUser, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, utils.BuildResponseSuccess(constants.SuccessUserUpdated, resp))
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	idParam := c.Param(constants.ParamID)
	if idParam == "" {
		c.JSON(http.StatusBadRequest, utils.BuildResponseFailed(constants.ErrInvalidRequest, constants.ErrIDRequired, nil))
		return
	}

	reqID := utils.StringToUint(idParam)
	if reqID == 0 {
		c.JSON(http.StatusBadRequest, utils.BuildResponseFailed(constants.ErrInvalidRequest, constants.ErrInvalidID, nil))
		return
	}

	grpcReq := &userpb.UserDeleteRequest{
		Id: uint64(reqID),
	}

	_, err := uc.UserClient.DeleteUser(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildResponseFailed(constants.ErrDeleteUser, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, utils.BuildResponseSuccess(constants.SuccessUserDeleted, nil))
}
