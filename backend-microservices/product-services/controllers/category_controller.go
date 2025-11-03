package controllers

import (
	"net/http"

	"product-services/constants"
	"product-services/dto"
	"product-services/services"
	"product-services/utils"

	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	GetAllCategories(c *gin.Context)
	CreateCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
}

type categoryController struct {
	categoryService services.CategoryService
}

func NewCategoryController(categoryService services.CategoryService) CategoryController {
	return &categoryController{categoryService: categoryService}
}

func (c *categoryController) GetAllCategories(ctx *gin.Context) {
	categories, err := c.categoryService.GetAllCategories(ctx)
	if err != nil {
		res := utils.BuildResponseFailed(constants.CATEGORY_NOT_FOUND, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.CATEGORY_RETRIEVED_SUCCESSFULLY, categories)
	ctx.JSON(http.StatusOK, res)
}

func (c *categoryController) CreateCategory(ctx *gin.Context) {
	var categoryRequest dto.CreateCategoryRequestDTO
	if err := ctx.ShouldBindJSON(&categoryRequest); err != nil {
		res := utils.BuildResponseFailed(constants.CATEGORY_CREATION_FAILED, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	category, err := c.categoryService.CreateCategory(ctx, categoryRequest)
	if err != nil {
		res := utils.BuildResponseFailed(constants.CATEGORY_CREATION_FAILED, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.CATEGORY_CREATED_SUCCESSFULLY, category)
	ctx.JSON(http.StatusCreated, res)
}

func (c *categoryController) UpdateCategory(ctx *gin.Context) {
	categoryIDParam := ctx.Param("id")
	if categoryIDParam == "" {
		res := utils.BuildResponseFailed(constants.CATEGORY_DELETION_FAILED, "category id is required", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	categoryID := utils.StringToUint(categoryIDParam)
	var categoryRequest dto.UpdateCategoryRequestDTO
	if err := ctx.ShouldBindJSON(&categoryRequest); err != nil {
		res := utils.BuildResponseFailed(constants.CATEGORY_UPDATE_FAILED, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	category, err := c.categoryService.UpdateCategory(ctx, categoryID, categoryRequest)
	if err != nil {
		res := utils.BuildResponseFailed(constants.CATEGORY_UPDATE_FAILED, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.CATEGORY_UPDATED_SUCCESSFULLY, category)
	ctx.JSON(http.StatusOK, res)
}

func (c *categoryController) DeleteCategory(ctx *gin.Context) {
	categoryIDParam := ctx.Param("id")
	if categoryIDParam == "" {
		res := utils.BuildResponseFailed(constants.CATEGORY_DELETION_FAILED, "category id is required", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	categoryID := utils.StringToUint(categoryIDParam)
	err := c.categoryService.DeleteCategory(ctx, categoryID)
	if err != nil {
		res := utils.BuildResponseFailed(constants.CATEGORY_DELETION_FAILED, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.CATEGORY_DELETED_SUCCESSFULLY, nil)
	ctx.JSON(http.StatusOK, res)
}
