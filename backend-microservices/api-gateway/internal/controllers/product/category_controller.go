package product

import (
	"net/http"

	constants "api-gateway/internal/constants"
	dto "api-gateway/internal/dto/product"
	productpb "api-gateway/internal/pb/product"
	"api-gateway/internal/utils"

	"github.com/gin-gonic/gin"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type CategoryController interface {
	GetAllCategories(c *gin.Context)
	CreateCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
}

type categoryController struct {
	productpb.UnimplementedCategoryServiceServer
	categoryClient productpb.CategoryServiceClient
}

func NewCategoryController(categoryClient productpb.CategoryServiceClient) CategoryController {
	return &categoryController{categoryClient: categoryClient}
}

func (c *categoryController) GetAllCategories(ctx *gin.Context) {
	categories, err := c.categoryClient.GetAllCategories(ctx, &emptypb.Empty{})
	if err != nil {
		res := utils.BuildResponseFailed(constants.ErrGetCategories, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.SuccessGetCategories, categories)
	ctx.JSON(http.StatusOK, res)
}

func (c *categoryController) CreateCategory(ctx *gin.Context) {
	var req dto.CreateCategoryRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(constants.ErrInvalidRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	gRPCReq := &productpb.CreateCategoryRequest{
		Name: req.Name,
	}

	category, err := c.categoryClient.CreateCategory(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ErrCreateCategory, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.SuccessCreateCategory, category)
	ctx.JSON(http.StatusCreated, res)
}

func (c *categoryController) UpdateCategory(ctx *gin.Context) {
	id := ctx.Param(constants.ParamID)
	if id == "" {
		res := utils.BuildResponseFailed(constants.ErrInvalidRequest, constants.ErrCategoryIDRequired, nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.UpdateCategoryRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(constants.ErrInvalidRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	categoryID := utils.StringToUint(id)
	gRPCReq := &productpb.UpdateCategoryRequest{
		Id:   categoryID,
		Name: req.Name,
	}

	category, err := c.categoryClient.UpdateCategory(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ErrUpdateCategory, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.SuccessUpdateCategory, category)
	ctx.JSON(http.StatusOK, res)
}

func (c *categoryController) DeleteCategory(ctx *gin.Context) {
	id := ctx.Param(constants.ParamID)
	if id == "" {
		res := utils.BuildResponseFailed(constants.ErrInvalidRequest, constants.ErrCategoryIDRequired, nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	categoryID := utils.StringToUint(id)
	gRPCReq := &productpb.DeleteCategoryRequest{Id: categoryID}

	_, err := c.categoryClient.DeleteCategory(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ErrDeleteCategory, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.SuccessDeleteCategory, nil)
	ctx.JSON(http.StatusOK, res)
}
