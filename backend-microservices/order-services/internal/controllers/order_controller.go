package controllers

import (
	"context"

	"order-services/internal/constants"
	"order-services/internal/dto"
	"order-services/internal/services"
	"order-services/internal/utils"
	orderpb "order-services/pb/order"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type OrderController struct {
	orderpb.UnimplementedOrderServiceServer
	orderService services.OrderService
}

func NewOrderController(orderService services.OrderService) *OrderController {
	return &OrderController{orderService: orderService}
}

func (c *OrderController) GetOrders(ctx context.Context, req *emptypb.Empty) (*orderpb.OrdersListResponse, error) {
	orders, err := c.orderService.GetOrders(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, constants.ErrOrderGet, err)
	}

	return orders, nil
}

func (c *OrderController) GetOrdersByUser(ctx context.Context, req *orderpb.GetOrdersByUserRequest) (*orderpb.OrdersListResponse, error) {
	orders, err := c.orderService.GetOrdersByUser(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, constants.ErrOrderGet, err)
	}

	return orders, nil
}

func (c *OrderController) GetOrdersByStatus(ctx context.Context, req *orderpb.GetOrdersByStatusRequest) (*orderpb.OrdersListResponse, error) {
	orders, err := c.orderService.GetOrdersByStatus(ctx, req.Status)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, constants.ErrOrderGet, err)
	}

	return orders, nil
}

func (c *OrderController) GetDetailOrder(ctx context.Context, req *orderpb.GetDetailOrderRequest) (*orderpb.OrderResponse, error) {
	order, err := c.orderService.GetDetailOrder(ctx, req.OrderId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, constants.ErrOrderGet, err)
	}

	return order, nil
}

func (c *OrderController) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.OrderResponse, error) {
	request := dto.CreateOrderRequestDto{
		UserID:        req.UserId,
		Total:         req.Total,
		PaymentMethod: req.PaymentMethod,
		Item:          utils.MapItemRPCToItemDTO(req.Item),
	}

	createdOrder, err := c.orderService.CreateOrder(ctx, request)
	if err != nil {
		return nil, status.Errorf(codes.Internal, constants.ErrOrderCreate, err)
	}

	return createdOrder, nil
}

func (c *OrderController) UpdateStatusOrder(ctx context.Context, req *orderpb.UpdateStatusOrderRequest) (*orderpb.OrderResponse, error) {
	request := dto.UpdateStatusOrderRequestDto{
		OrderID: req.OrderId,
		Status:  req.Status,
	}

	updatedOrder, err := c.orderService.UpdateStatusOrder(ctx, request)
	if err != nil {
		return nil, status.Errorf(codes.Internal, constants.ErrOrderUpdate, err)
	}

	return updatedOrder, nil
}

func (c *OrderController) DeleteOrder(ctx context.Context, req *orderpb.DeleteOrderRequest) (*emptypb.Empty, error) {
	err := c.orderService.DeleteOrder(ctx, req.OrderId, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, constants.ErrOrderDelete, err)
	}

	return &emptypb.Empty{}, nil
}
