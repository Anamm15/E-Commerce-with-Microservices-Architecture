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

type ReviewController interface {
	GetAllReviews(c *gin.Context)
	GetReviewByProductID(c *gin.Context)
	CreateReview(c *gin.Context)
	UpdateReview(c *gin.Context)
	DeleteReview(c *gin.Context)
}

type reviewController struct {
	productpb.UnimplementedReviewServiceServer
	reviewClient productpb.ReviewServiceClient
}

func NewReviewController(reviewClient productpb.ReviewServiceClient) ReviewController {
	return &reviewController{reviewClient: reviewClient}
}

func (c *reviewController) GetAllReviews(ctx *gin.Context) {
	reviews, err := c.reviewClient.GetAllReviews(ctx, &emptypb.Empty{})
	if err != nil {
		res := utils.BuildResponseFailed("Failed to get reviews", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Reviews retrieved successfully", reviews)
	ctx.JSON(http.StatusOK, res)
}

func (c *reviewController) GetReviewByProductID(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		res := utils.BuildResponseFailed("Invalid request", "product id is required", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	productID, _ := strconv.ParseUint(id, 10, 32)

	gRPCReq := &productpb.GetReviewByProductRequest{
		ProductId: uint32(productID),
	}

	reviews, err := c.reviewClient.GetReviewByProductID(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to get reviews by product id", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Reviews retrieved successfully", reviews)
	ctx.JSON(http.StatusOK, res)
}

func (c *reviewController) CreateReview(ctx *gin.Context) {
	var req dto.CreateReviewRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	gRPCReq := &productpb.CreateReviewRequest{
		ProductId: uint32(req.ProductID),
		UserId:    uint32(req.UserID),
		Rating:    int32(req.Rating),
		Comment:   req.Comment,
	}

	review, err := c.reviewClient.CreateReview(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to create review", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Review created successfully", review)
	ctx.JSON(http.StatusCreated, res)
}

func (c *reviewController) UpdateReview(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		res := utils.BuildResponseFailed("Invalid request", "review id is required", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	reviewID, _ := strconv.ParseUint(id, 10, 64)

	var req dto.UpdateReviewRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("Invalid request", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	gRPCReq := &productpb.UpdateReviewRequest{
		ReviewId:  uint32(reviewID),
		ProductId: uint32(req.ProductID),
		Rating:    int32(req.Rating),
		Comment:   req.Comment,
		UserId:    uint32(req.UserID),
	}

	review, err := c.reviewClient.UpdateReview(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to update review", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Review updated successfully", review)
	ctx.JSON(http.StatusOK, res)
}

func (c *reviewController) DeleteReview(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		res := utils.BuildResponseFailed("Invalid request", "review id is required", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	reviewID, _ := strconv.ParseUint(id, 10, 32)

	gRPCReq := &productpb.DeleteReviewRequest{ReviewId: uint32(reviewID)}

	_, err := c.reviewClient.DeleteReview(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to delete review", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Review deleted successfully", nil)
	ctx.JSON(http.StatusOK, res)
}
