package controllers

import (
	"net/http"

	"user-services/dto"
	"user-services/services"

	"user-services/utils"

	"user-services/constants"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetAllUser(ctx *gin.Context)
	GetUserByUsername(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (c *userController) GetAllUser(ctx *gin.Context) {
	users, err := c.userService.GetAllUser(ctx)
	if err != nil {
		res := utils.BuildResponseFailed(constants.USER_NOT_FOUND, err.Error(), ctx)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.USER_RETRIEVED_SUCCESSFULLY, users)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) GetUserByUsername(ctx *gin.Context) {
	usernameParam := ctx.Query("username")
	user, err := c.userService.GetUserByUsername(ctx, usernameParam)
	if err != nil {
		res := utils.BuildResponseFailed(constants.USER_NOT_FOUND, err.Error(), ctx)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.USER_RETRIEVED_SUCCESSFULLY, user)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) CreateUser(ctx *gin.Context) {
	var user dto.UserCreateDTO
	err := ctx.ShouldBind(&user)
	if err != nil {
		res := utils.BuildResponseFailed(constants.INVALID_REQUEST, err.Error(), ctx)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	createdUser, err := c.userService.CreateUser(ctx, user)
	if err != nil {
		res := utils.BuildResponseFailed(constants.USER_CREATION_FAILED, err.Error(), ctx)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.USER_CREATED_SUCCESSFULLY, createdUser)
	ctx.JSON(http.StatusCreated, res)
}

func (c *userController) UpdateUser(ctx *gin.Context) {
	var user dto.UserUpdateDTO
	err := ctx.ShouldBind(&user)
	if err != nil {
		res := utils.BuildResponseFailed(constants.INVALID_REQUEST, err.Error(), ctx)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	updatedUser, err := c.userService.UpdateUser(ctx, user)
	if err != nil {
		res := utils.BuildResponseFailed(constants.USER_UPDATE_FAILED, err.Error(), ctx)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.USER_UPDATED_SUCCESSFULLY, updatedUser)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) DeleteUser(ctx *gin.Context) {
	var user dto.UserDeleteDTO
	err := ctx.ShouldBind(&user)
	if err != nil {
		res := utils.BuildResponseFailed(constants.INVALID_REQUEST, err.Error(), ctx)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err = c.userService.DeleteUser(ctx, user)
	if err != nil {
		res := utils.BuildResponseFailed(constants.USER_DELETION_FAILED, err.Error(), ctx)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.USER_DELETED_SUCCESSFULLY, nil)
	ctx.JSON(http.StatusOK, res)
}
