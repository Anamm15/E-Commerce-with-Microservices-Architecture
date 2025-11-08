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

type ProductController interface {
	GetAllProducts(ctx *gin.Context)
	GetProductById(ctx *gin.Context)
	GetProductByCategoryID(ctx *gin.Context)
	CreateProduct(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
}

type productController struct {
	productpb.UnimplementedProductServiceServer
	productClient productpb.ProductServiceClient
}

func NewProductController(productClient productpb.ProductServiceClient) ProductController {
	return &productController{productClient: productClient}
}

func (c *productController) GetAllProducts(ctx *gin.Context) {
	products, err := c.productClient.GetAllProducts(ctx, &emptypb.Empty{})
	if err != nil {
		res := utils.BuildResponseFailed(constants.ErrGetProducts, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.SuccessGetProducts, products.Products)
	ctx.JSON(http.StatusOK, res)
}

func (c *productController) GetProductById(ctx *gin.Context) {
	id := ctx.Param(constants.ParamID)
	if id == "" {
		res := utils.BuildResponseFailed(constants.ErrInvalidRequest, constants.ErrProductIDRequired, nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	productID := utils.StringToUint(id)
	gRPCReq := &productpb.GetProductByIDRequest{Id: productID}

	product, err := c.productClient.GetProductByID(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ErrGetProduct, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.SuccessGetProduct, product.Product)
	ctx.JSON(http.StatusOK, res)
}

func (c *productController) GetProductByCategoryID(ctx *gin.Context) {
	id := ctx.Param(constants.ParamID)
	if id == "" {
		res := utils.BuildResponseFailed(constants.ErrInvalidRequest, constants.ErrCategoryIDRequired, nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	categoryID := utils.StringToUint(id)
	gRPCReq := &productpb.GetProductByCategoryRequest{CategoryId: categoryID}

	products, err := c.productClient.GetProductByCategoryID(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ErrGetProductsByCategory, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.SuccessGetProducts, products.Products)
	ctx.JSON(http.StatusOK, res)
}

func (c *productController) CreateProduct(ctx *gin.Context) {
	var req dto.CreateProductRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(constants.ErrInvalidRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	gRPCReq := &productpb.CreateProductRequest{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Category:    req.Category,
	}

	product, err := c.productClient.CreateProduct(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ErrCreateProduct, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.SuccessCreateProduct, product.Product)
	ctx.JSON(http.StatusCreated, res)
}

func (c *productController) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param(constants.ParamID)
	if id == "" {
		res := utils.BuildResponseFailed(constants.ErrInvalidRequest, constants.ErrProductIDRequired, nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.UpdateProductRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(constants.ErrInvalidRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	productID := utils.StringToUint(id)
	gRPCReq := &productpb.UpdateProductRequest{
		Id:          productID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Category:    req.Category,
	}

	product, err := c.productClient.UpdateProduct(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ErrUpdateProduct, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.SuccessUpdateProduct, product.Product)
	ctx.JSON(http.StatusOK, res)
}

func (c *productController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param(constants.ParamID)
	if id == "" {
		res := utils.BuildResponseFailed(constants.ErrInvalidRequest, constants.ErrProductIDRequired, nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	productID := utils.StringToUint(id)
	gRPCReq := &productpb.DeleteProductRequest{Id: productID}

	_, err := c.productClient.DeleteProduct(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ErrDeleteProduct, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.SuccessDeleteProduct, nil)
	ctx.JSON(http.StatusOK, res)
}
