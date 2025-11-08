package product

import (
	"net/http"
	"strconv"

	dto "api-gateway/internal/dto/product"
	productpb "api-gateway/internal/pb/product"
	"api-gateway/internal/utils"

	emptypb "google.golang.org/protobuf/types/known/emptypb"

	"github.com/gin-gonic/gin"
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
		res := utils.BuildResponseFailed("Failed to get products", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Products retrieved successfully", products.Products)
	ctx.JSON(http.StatusOK, res)
}

func (c *productController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		res := utils.BuildResponseFailed("Invalid request", "product id is required", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	productID, _ := strconv.ParseUint(id, 10, 64)
	gRPCReq := &productpb.GetProductByIDRequest{Id: uint32(productID)}

	product, err := c.productClient.GetProductByID(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to get product", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Product retrieved successfully", product.Product)
	ctx.JSON(http.StatusOK, res)
}

func (c *productController) GetProductByCategoryID(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		res := utils.BuildResponseFailed("Invalid request", "category id is required", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	categoryID, _ := strconv.ParseUint(id, 10, 32)
	gRPCReq := &productpb.GetProductByCategoryRequest{CategoryId: uint32(categoryID)}

	products, err := c.productClient.GetProductByCategoryID(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to get products by category", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Products retrieved successfully", products.Products)
	ctx.JSON(http.StatusOK, res)
}

func (c *productController) CreateProduct(ctx *gin.Context) {
	var req dto.CreateProductRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	gRPCReq := &productpb.CreateProductRequest{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       int32(req.Stock),
		Category:    req.Category,
	}

	product, err := c.productClient.CreateProduct(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to create product", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Product created successfully", product.Product)
	ctx.JSON(http.StatusCreated, res)
}

func (c *productController) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		res := utils.BuildResponseFailed("Invalid request", "product id is required", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	productID, _ := strconv.ParseUint(id, 10, 32)
	var req dto.UpdateProductRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	gRPCReq := &productpb.UpdateProductRequest{
		Id:          uint32(productID),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       int32(req.Stock),
		Category:    req.Category,
	}

	product, err := c.productClient.UpdateProduct(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to update product", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Product updated successfully", product.Product)
	ctx.JSON(http.StatusOK, res)
}

func (c *productController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		res := utils.BuildResponseFailed("Invalid request", "product id is required", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	productID, _ := strconv.ParseUint(id, 10, 32)
	gRPCReq := &productpb.DeleteProductRequest{Id: uint32(productID)}

	_, err := c.productClient.DeleteProduct(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to delete product", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Product deleted successfully", nil)
	ctx.JSON(http.StatusOK, res)
}
