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
		res := utils.BuildResponseFailed(constants.ErrGetReviews, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.SuccessGetReviews, reviews)
	ctx.JSON(http.StatusOK, res)
}

func (c *reviewController) GetReviewByProductID(ctx *gin.Context) {
	id := ctx.Param(constants.ParamID)
	if id == "" {
		res := utils.BuildResponseFailed(constants.ErrInvalidRequest, constants.ErrProductIDRequired, nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	productID := utils.StringToUint(id)

	gRPCReq := &productpb.GetReviewByProductRequest{
		ProductId: productID,
	}

	reviews, err := c.reviewClient.GetReviewByProductID(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ErrGetReviewsByProductID, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.SuccessGetReviews, reviews)
	ctx.JSON(http.StatusOK, res)
}

func (c *reviewController) CreateReview(ctx *gin.Context) {
	userID := ctx.MustGet(constants.ContextKeyUserID).(uint64)
	var req dto.CreateReviewRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(constants.ErrInvalidRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	gRPCReq := &productpb.CreateReviewRequest{
		ProductId: req.ProductID,
		UserId:    userID,
		Rating:    req.Rating,
		Comment:   req.Comment,
	}

	review, err := c.reviewClient.CreateReview(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ErrCreateReview, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.SuccessCreateReview, review)
	ctx.JSON(http.StatusCreated, res)
}

func (c *reviewController) UpdateReview(ctx *gin.Context) {
	id := ctx.Param(constants.ParamID)
	userID := ctx.MustGet(constants.ContextKeyUserID).(uint64)
	if id == "" || userID == 0 {
		res := utils.BuildResponseFailed(constants.ErrInvalidRequest, constants.ErrReviewIDOrLoginRequired, nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.UpdateReviewRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(constants.ErrInvalidRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reviewID := utils.StringToUint(id)
	gRPCReq := &productpb.UpdateReviewRequest{
		ReviewId:  reviewID,
		ProductId: req.ProductID,
		Rating:    req.Rating,
		Comment:   req.Comment,
		UserId:    userID,
	}

	review, err := c.reviewClient.UpdateReview(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ErrUpdateReview, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.SuccessUpdateReview, review)
	ctx.JSON(http.StatusOK, res)
}

func (c *reviewController) DeleteReview(ctx *gin.Context) {
	id := ctx.Param(constants.ParamID)
	userID := ctx.MustGet(constants.ContextKeyUserID).(uint64)
	if id == "" {
		res := utils.BuildResponseFailed(constants.ErrInvalidRequest, constants.ErrReviewIDRequired, nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reviewID := utils.StringToUint(id)
	gRPCReq := &productpb.DeleteReviewRequest{ReviewId: reviewID, UserId: userID}

	_, err := c.reviewClient.DeleteReview(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ErrDeleteReview, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.SuccessDeleteReview, nil)
	ctx.JSON(http.StatusOK, res)
}
