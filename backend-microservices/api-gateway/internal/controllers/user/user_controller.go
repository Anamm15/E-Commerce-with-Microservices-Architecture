package user

import (
	"context"
	"net/http"

	dto "api-gateway/internal/dto/user"
	userpb "api-gateway/internal/pb/user"
	"api-gateway/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserClient userpb.UserServiceClient
}

func NewUserController(UserClient userpb.UserServiceClient) *UserController {
	return &UserController{UserClient: UserClient}
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	resp, err := uc.UserClient.GetAllUsers(context.Background(), &userpb.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildResponseFailed("failed to get users", err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, utils.BuildResponseSuccess("users fetched", resp.Users))
}

func (uc *UserController) GetUserByUsername(c *gin.Context) {
	username := c.Query("username")

	grpcReq := &userpb.GetUserByUsernameRequest{Username: username}

	resp, err := uc.UserClient.GetUserByUsername(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildResponseFailed("failed to get user", err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, utils.BuildResponseSuccess("user fetched", resp))
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var req dto.UserCreateDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildResponseFailed("invalid request", err.Error(), nil))
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
		c.JSON(http.StatusInternalServerError, utils.BuildResponseFailed("gRPC error", err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, utils.BuildResponseSuccess("user created", resp))
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, utils.BuildResponseFailed("invalid request", "id is required", nil))
		return
	}

	reqID := utils.StringToUint(idParam)
	if reqID == 0 {
		c.JSON(http.StatusBadRequest, utils.BuildResponseFailed("invalid request", "invalid id", nil))
		return
	}

	var req dto.UserUpdateDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildResponseFailed("invalid request", err.Error(), nil))
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
		c.JSON(http.StatusInternalServerError, utils.BuildResponseFailed("failed to update user", err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, utils.BuildResponseSuccess("user updated", resp))
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, utils.BuildResponseFailed("invalid request", "id is required", nil))
		return
	}

	reqID := utils.StringToUint(idParam)
	if reqID == 0 {
		c.JSON(http.StatusBadRequest, utils.BuildResponseFailed("invalid request", "invalid id", nil))
		return
	}

	grpcReq := &userpb.UserDeleteRequest{
		Id: uint64(reqID),
	}

	_, err := uc.UserClient.DeleteUser(context.Background(), grpcReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildResponseFailed("failed to delete user", err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, utils.BuildResponseSuccess("user deleted", nil))
}
