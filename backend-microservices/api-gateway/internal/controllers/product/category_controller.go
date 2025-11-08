package product

import (
	"net/http"
	"strconv"

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
		res := utils.BuildResponseFailed("Failed to get categories", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Categories retrieved successfully", categories)
	ctx.JSON(http.StatusOK, res)
}

func (c *categoryController) CreateCategory(ctx *gin.Context) {
	var req dto.CreateCategoryRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	gRPCReq := &productpb.CreateCategoryRequest{
		Name: req.Name,
	}

	category, err := c.categoryClient.CreateCategory(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to create category", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Category created successfully", category)
	ctx.JSON(http.StatusCreated, res)
}

func (c *categoryController) UpdateCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		res := utils.BuildResponseFailed("Invalid request", "category id is required", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	categoryID, _ := strconv.ParseUint(id, 10, 64)

	var req dto.UpdateCategoryRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	gRPCReq := &productpb.UpdateCategoryRequest{
		Id:   uint32(categoryID),
		Name: req.Name,
	}

	category, err := c.categoryClient.UpdateCategory(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to update category", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Category updated successfully", category)
	ctx.JSON(http.StatusOK, res)
}

func (c *categoryController) DeleteCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		res := utils.BuildResponseFailed("Invalid request", "category id is required", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	categoryID, _ := strconv.ParseUint(id, 10, 32)

	gRPCReq := &productpb.DeleteCategoryRequest{Id: uint32(categoryID)}

	_, err := c.categoryClient.DeleteCategory(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to delete category", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Category deleted successfully", nil)
	ctx.JSON(http.StatusOK, res)
}
