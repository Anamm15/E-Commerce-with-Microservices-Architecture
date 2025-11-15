package order

import (
	"net/http"

	constants "api-gateway/internal/constants"
	dto "api-gateway/internal/dto/order"
	orderpb "api-gateway/internal/pb/order"
	"api-gateway/internal/utils"

	"github.com/gin-gonic/gin"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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
	orderpb.UnimplementedOrderServiceServer
	orderClient orderpb.OrderServiceClient
}

func NewOrderController(orderClient orderpb.OrderServiceClient) OrderController {
	return &orderController{orderClient: orderClient}
}

func (c *orderController) GetOrders(ctx *gin.Context) {
	result, err := c.orderClient.GetOrders(ctx, &emptypb.Empty{})
	if err != nil {
		res := utils.BuildResponseFailed(constants.ORDER_NOT_FOUND, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.ORDER_RETRIEVED_SUCCESSFULLY, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *orderController) GetOrdersByUser(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(uint64)

	gRPCReq := &orderpb.GetOrdersByUserRequest{UserId: userId}
	resp, err := c.orderClient.GetOrdersByUser(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ORDER_NOT_FOUND, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.ORDER_RETRIEVED_SUCCESSFULLY, resp)
	ctx.JSON(http.StatusOK, res)
}

func (c *orderController) GetOrdersByStatus(ctx *gin.Context) {
	status := ctx.Query("status")

	gRPCReq := &orderpb.GetOrdersByStatusRequest{Status: status}
	resp, err := c.orderClient.GetOrdersByStatus(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ORDER_NOT_FOUND, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.ORDER_RETRIEVED_SUCCESSFULLY, resp)
	ctx.JSON(http.StatusOK, res)
}

func (c *orderController) GetDetailOrder(ctx *gin.Context) {
	orderIdParams := ctx.Param("id")
	orderId := utils.StringToUint(orderIdParams)

	gRPCReq := &orderpb.GetDetailOrderRequest{OrderId: orderId}
	resp, err := c.orderClient.GetDetailOrder(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ORDER_NOT_FOUND, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.ORDER_RETRIEVED_SUCCESSFULLY, resp)
	ctx.JSON(http.StatusOK, res)
}

func (c *orderController) CreateOrder(ctx *gin.Context) {
	var req dto.CreateOrderRequestDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(constants.ORDER_CREATION_FAILED, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	userId := ctx.MustGet("user_id").(uint64)

	var items []*orderpb.CreateOrderItem
	for _, it := range req.Item {
		items = append(items, &orderpb.CreateOrderItem{
			ProductId: it.ProductID,
			Quantity:  it.Quantity,
			Total:     it.Total,
		})
	}

	gRPCReq := &orderpb.CreateOrderRequest{
		UserId:        userId,
		Total:         req.Total,
		PaymentMethod: req.PaymentMethod,
		Item:          items,
	}

	result, err := c.orderClient.CreateOrder(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ORDER_CREATION_FAILED, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.ORDER_CREATED_SUCCESSFULLY, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *orderController) UpdateStatusOrder(ctx *gin.Context) {
	var req dto.UpdateStatusOrderRequestDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(constants.ORDER_UPDATE_FAILED, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	gRPCReq := &orderpb.UpdateStatusOrderRequest{
		OrderId: req.OrderID,
		Status:  req.Status,
	}

	result, err := c.orderClient.UpdateStatusOrder(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ORDER_UPDATE_FAILED, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.ORDER_UPDATED_SUCCESSFULLY, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *orderController) DeleteOrder(ctx *gin.Context) {
	orderIDParams := ctx.Param("id")
	orderID := utils.StringToUint(orderIDParams)
	userID := ctx.MustGet("user_id").(uint64)

	gRPCReq := &orderpb.DeleteOrderRequest{
		OrderId: orderID,
		UserId:  userID,
	}

	_, err := c.orderClient.DeleteOrder(ctx, gRPCReq)
	if err != nil {
		res := utils.BuildResponseFailed(constants.ORDER_DELETION_FAILED, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(constants.ORDER_DELETED_SUCCESSFULLY, nil)
	ctx.JSON(http.StatusOK, res)
}
