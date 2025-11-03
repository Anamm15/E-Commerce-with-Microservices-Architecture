package controllers

import (
	"net/http"

	"product-services/constants"
	"product-services/dto"
	"product-services/services"
	"product-services/utils"

	"github.com/gin-gonic/gin"
)

type ProductController interface {
	GetAllProducts(c *gin.Context)
	GetProductById(c *gin.Context)
	GetProductByCategoryID(c *gin.Context)
	CreateProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}

type productController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) ProductController {
	return &productController{productService: productService}
}

func (c *productController) GetAllProducts(ctx *gin.Context) {
	products, err := c.productService.GetAllProducts(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed(constants.PRODUCT_NOT_FOUND, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.PRODUCT_RETRIEVED_SUCCESSFULLY, products)
	ctx.JSON(http.StatusOK, res)
}

func (c *productController) GetProductById(ctx *gin.Context) {
	productIDParam := ctx.Param("id")
	if productIDParam == "" {
		res := utils.BuildResponseFailed(constants.PRODUCT_NOT_FOUND, "product id is required", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	productID := utils.StringToUint(productIDParam)
	product, err := c.productService.GetProductByID(ctx.Request.Context(), productID)
	if err != nil {
		res := utils.BuildResponseFailed(constants.PRODUCT_NOT_FOUND, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.PRODUCT_RETRIEVED_SUCCESSFULLY, product)
	ctx.JSON(http.StatusOK, res)
}

func (c *productController) GetProductByCategoryID(ctx *gin.Context) {
	categoryIDParam := ctx.Param("id")
	if categoryIDParam == "" {
		res := utils.BuildResponseFailed(constants.PRODUCT_NOT_FOUND, "category id is required", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	categoryID := utils.StringToUint(categoryIDParam)
	products, err := c.productService.GetProductByCategoryID(ctx.Request.Context(), categoryID)
	if err != nil {
		res := utils.BuildResponseFailed(constants.PRODUCT_NOT_FOUND, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.PRODUCT_RETRIEVED_SUCCESSFULLY, products)
	ctx.JSON(http.StatusOK, res)
}

func (c *productController) CreateProduct(ctx *gin.Context) {
	var productRequest dto.CreateProductRequestDTO
	if err := ctx.ShouldBind(&productRequest); err != nil {
		res := utils.BuildResponseFailed(constants.PRODUCT_CREATION_FAILED, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	form, err := ctx.MultipartForm()
	if err == nil && form.File["images"] != nil {
		productRequest.Images = form.File["images"]
	}

	product, err := c.productService.CreateProduct(ctx.Request.Context(), productRequest)
	if err != nil {
		res := utils.BuildResponseFailed(constants.PRODUCT_CREATION_FAILED, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.PRODUCT_CREATED_SUCCESSFULLY, product)
	ctx.JSON(http.StatusCreated, res)
}

func (c *productController) UpdateProduct(ctx *gin.Context) {
	productIDParam := ctx.Param("id")
	if productIDParam == "" {
		res := utils.BuildResponseFailed(constants.PRODUCT_UPDATE_FAILED, "product id is required", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	productID := utils.StringToUint(productIDParam)
	var productRequest dto.UpdateProductRequestDTO
	if err := ctx.ShouldBindJSON(&productRequest); err != nil {
		res := utils.BuildResponseFailed(constants.PRODUCT_UPDATE_FAILED, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	product, err := c.productService.UpdateProduct(ctx.Request.Context(), productID, productRequest)
	if err != nil {
		res := utils.BuildResponseFailed(constants.PRODUCT_UPDATE_FAILED, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.PRODUCT_UPDATED_SUCCESSFULLY, product)
	ctx.JSON(http.StatusOK, res)
}

func (c *productController) DeleteProduct(ctx *gin.Context) {
	productIDParam := ctx.Param("id")
	if productIDParam == "" {
		res := utils.BuildResponseFailed(constants.PRODUCT_DELETION_FAILED, "product id is required", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	productID := utils.StringToUint(productIDParam)
	err := c.productService.DeleteProduct(ctx.Request.Context(), productID)
	if err != nil {
		res := utils.BuildResponseFailed(constants.PRODUCT_DELETION_FAILED, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.PRODUCT_DELETED_SUCCESSFULLY, nil)
	ctx.JSON(http.StatusOK, res)
}
