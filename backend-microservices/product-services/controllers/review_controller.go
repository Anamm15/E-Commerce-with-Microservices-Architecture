package controllers

import (
	"net/http"

	"product-services/constants"
	"product-services/dto"
	"product-services/services"
	"product-services/utils"

	"github.com/gin-gonic/gin"
)

type ReviewController interface {
	GetAllReviews(c *gin.Context)
	GetReviewByProductID(c *gin.Context)
	CreateReview(c *gin.Context)
	UpdateReview(c *gin.Context)
	DeleteReview(c *gin.Context)
}

type reviewController struct {
	reviewService services.ReviewService
}

func NewReviewController(reviewService services.ReviewService) ReviewController {
	return &reviewController{reviewService: reviewService}
}

func (c *reviewController) GetAllReviews(ctx *gin.Context) {
	reviews, err := c.reviewService.GetAllReviews(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed(constants.REVIEW_NOT_FOUND, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.REVIEW_RETRIEVED_SUCCESSFULLY, reviews)
	ctx.JSON(http.StatusOK, res)
}

func (c *reviewController) GetReviewByProductID(ctx *gin.Context) {
	productIDParam := ctx.Param("id")
	if productIDParam == "" {
		res := utils.BuildResponseFailed(constants.REVIEW_NOT_FOUND, "product id is required", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	productID := utils.StringToUint(productIDParam)
	reviews, err := c.reviewService.GetReviewProductID(ctx, productID)
	if err != nil {
		res := utils.BuildResponseFailed(constants.REVIEW_NOT_FOUND, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.REVIEW_RETRIEVED_SUCCESSFULLY, reviews)
	ctx.JSON(http.StatusOK, res)
}

func (c *reviewController) CreateReview(ctx *gin.Context) {
	var reviewRequest dto.CreateReviewRequestDTO
	if err := ctx.ShouldBindJSON(&reviewRequest); err != nil {
		res := utils.BuildResponseFailed(constants.REVIEW_CREATION_FAILED, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	review, err := c.reviewService.CreateReview(ctx, reviewRequest)
	if err != nil {
		res := utils.BuildResponseFailed(constants.REVIEW_CREATION_FAILED, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.REVIEW_CREATED_SUCCESSFULLY, review)
	ctx.JSON(http.StatusCreated, res)
}

func (c *reviewController) UpdateReview(ctx *gin.Context) {
	reviewIDParam := ctx.Param("id")
	if reviewIDParam == "" {
		res := utils.BuildResponseFailed(constants.REVIEW_UPDATE_FAILED, "review id is required", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reviewID := utils.StringToUint(reviewIDParam)
	var reviewRequest dto.UpdateReviewRequestDTO
	if err := ctx.ShouldBindJSON(&reviewRequest); err != nil {
		res := utils.BuildResponseFailed(constants.REVIEW_UPDATE_FAILED, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	review, err := c.reviewService.UpdateReview(ctx, reviewID, reviewRequest)
	if err != nil {
		res := utils.BuildResponseFailed(constants.REVIEW_UPDATE_FAILED, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.REVIEW_UPDATED_SUCCESSFULLY, review)
	ctx.JSON(http.StatusOK, res)
}

func (c *reviewController) DeleteReview(ctx *gin.Context) {
	reviewIDParam := ctx.Param("id")
	if reviewIDParam == "" {
		res := utils.BuildResponseFailed(constants.REVIEW_DELETION_FAILED, "review id is required", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reviewID := utils.StringToUint(reviewIDParam)
	err := c.reviewService.DeleteReview(ctx, reviewID)
	if err != nil {
		res := utils.BuildResponseFailed(constants.REVIEW_DELETION_FAILED, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.REVIEW_DELETED_SUCCESSFULLY, nil)
	ctx.JSON(http.StatusOK, res)
}
