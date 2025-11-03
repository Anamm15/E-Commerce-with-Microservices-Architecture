package controllers

import (
	"net/http"

	"order-services/constants"
	"order-services/dto"
	"order-services/services"
	"order-services/utils"

	"github.com/gin-gonic/gin"
)

type OrderController interface {
	GetOrders(ctx *gin.Context)
	GetOrdersByUser(ctx *gin.Context)
	GetOrdersByStatus(ctx *gin.Context)
	GetDetailOrder(ctx *gin.Context)
	CreateOrder(ctx *gin.Context)
	UpdateStatusOrder(ctx *gin.Context)
	DeleteOrder(ctx *gin.Context)
}

type orderController struct {
	orderService services.OrderService
}

func NewOrderController(orderService services.OrderService) OrderController {
	return &orderController{orderService: orderService}
}

func (c *orderController) GetOrders(ctx *gin.Context) {
	result, err := c.orderService.GetOrders(ctx)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ORDER_NOT_FOUND, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.ORDER_RETRIEVED_SUCCESSFULLY, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *orderController) GetOrdersByUser(ctx *gin.Context) {
	var orderId uint
	orderIdParams := ctx.Param("id")
	if orderIdParams != "" {
		orderId, _ = utils.StringToUint(orderIdParams)
	}

	result, err := c.orderService.GetOrdersByUser(ctx, orderId)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ORDER_NOT_FOUND, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.ORDER_RETRIEVED_SUCCESSFULLY, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *orderController) GetOrdersByStatus(ctx *gin.Context) {
	result, err := c.orderService.GetOrdersByStatus(ctx, ctx.Param("status"))
	if err != nil {
		res := utils.BuildResponseFailed(constants.ORDER_NOT_FOUND, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.ORDER_RETRIEVED_SUCCESSFULLY, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *orderController) GetDetailOrder(ctx *gin.Context) {
	var orderId uint
	orderIdParams := ctx.Param("id")
	if orderIdParams != "" {
		orderId, _ = utils.StringToUint(orderIdParams)
	}

	result, err := c.orderService.GetDetailOrder(ctx, orderId)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ORDER_NOT_FOUND, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.ORDER_RETRIEVED_SUCCESSFULLY, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *orderController) CreateOrder(ctx *gin.Context) {
	var request dto.CreateOrderRequestDto
	if err := ctx.ShouldBindJSON(&request); err != nil {
		res := utils.BuildResponseFailed(constants.ORDER_CREATION_FAILED, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.orderService.CreateOrder(ctx, request)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ORDER_CREATION_FAILED, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.ORDER_CREATED_SUCCESSFULLY, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *orderController) UpdateStatusOrder(ctx *gin.Context) {
	var request dto.UpdateStatusOrderRequestDto
	if err := ctx.ShouldBindJSON(&request); err != nil {
		res := utils.BuildResponseFailed(constants.ORDER_UPDATE_FAILED, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.orderService.UpdateStatusOrder(ctx, request)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ORDER_UPDATE_FAILED, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.ORDER_UPDATED_SUCCESSFULLY, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *orderController) DeleteOrder(ctx *gin.Context) {
	var orderId uint
	orderIdParams := ctx.Param("id")
	if orderIdParams != "" {
		orderId, _ = utils.StringToUint(orderIdParams)
	}

	err := c.orderService.DeleteOrder(ctx, orderId)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ORDER_DELETION_FAILED, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.ORDER_DELETED_SUCCESSFULLY, nil)
	ctx.JSON(http.StatusOK, res)
}
