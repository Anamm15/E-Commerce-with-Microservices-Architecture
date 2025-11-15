package services

import (
	"context"

	"order-services/internal/dto"
	"order-services/internal/repositories"
	"order-services/internal/utils"
	orderpb "order-services/pb/order"
	productpb "order-services/pb/product"
)

type OrderService interface {
	GetOrders(ctx context.Context) (*orderpb.OrdersListResponse, error)
	GetOrdersByUser(ctx context.Context, userID uint64) (*orderpb.OrdersListResponse, error)
	GetOrdersByStatus(ctx context.Context, status string) (*orderpb.OrdersListResponse, error)
	GetDetailOrder(ctx context.Context, orderID uint64) (*orderpb.OrderResponse, error)
	CreateOrder(ctx context.Context, order dto.CreateOrderRequestDto) (*orderpb.OrderResponse, error)
	UpdateStatusOrder(ctx context.Context, order dto.UpdateStatusOrderRequestDto) (*orderpb.OrderResponse, error)
	DeleteOrder(ctx context.Context, orderId uint64, userId uint64) error
}

type orderService struct {
	orderRepository     repositories.OrderRepository
	orderItemRepository repositories.OrderItemRepository
	productClient       productpb.ProductServiceClient
}

func NewOrderService(
	orderRepository repositories.OrderRepository,
	orderItemRepository repositories.OrderItemRepository,
	productClient productpb.ProductServiceClient,
) OrderService {
	return &orderService{
		orderRepository:     orderRepository,
		orderItemRepository: orderItemRepository,
		productClient:       productClient,
	}
}

func (s *orderService) GetOrders(ctx context.Context) (*orderpb.OrdersListResponse, error) {
	orders, err := s.orderRepository.GetOrders(ctx)
	if err != nil {
		return nil, err
	}

	ordersResponseDTO := utils.MapOrdersModelToDTO(orders)
	ordersWithProduct, err := utils.AttachProductDetailToOrders(ctx, ordersResponseDTO, s.productClient)
	if err != nil {
		return nil, err
	}

	res := utils.MapOrderResponseDTOToOrderListRPC(ordersWithProduct)
	return &orderpb.OrdersListResponse{
		Orders: res,
	}, nil
}

func (s *orderService) GetOrdersByUser(ctx context.Context, userID uint64) (*orderpb.OrdersListResponse, error) {
	orders, err := s.orderRepository.GetOrdersByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	ordersResponseDTO := utils.MapOrdersModelToDTO(orders)
	ordersWithProduct, err := utils.AttachProductDetailToOrders(ctx, ordersResponseDTO, s.productClient)
	if err != nil {
		return nil, err
	}

	res := utils.MapOrderResponseDTOToOrderListRPC(ordersWithProduct)
	return &orderpb.OrdersListResponse{
		Orders: res,
	}, nil
}

func (s *orderService) GetOrdersByStatus(ctx context.Context, status string) (*orderpb.OrdersListResponse, error) {
	orders, err := s.orderRepository.GetOrdersByStatus(ctx, status)
	if err != nil {
		return nil, err
	}

	ordersResponseDTO := utils.MapOrdersModelToDTO(orders)
	ordersWithProduct, err := utils.AttachProductDetailToOrders(ctx, ordersResponseDTO, s.productClient)
	if err != nil {
		return nil, err
	}

	res := utils.MapOrderResponseDTOToOrderListRPC(ordersWithProduct)
	return &orderpb.OrdersListResponse{
		Orders: res,
	}, nil
}

func (s *orderService) GetDetailOrder(ctx context.Context, orderID uint64) (*orderpb.OrderResponse, error) {
	order, err := s.orderRepository.GetDetailOrder(ctx, orderID)
	if err != nil {
		return nil, err
	}

	orderResponseDTO := utils.MapOrderModelToDTO(order)
	ordersWithProduct, err := utils.AttachProductDetailToOrder(ctx, orderResponseDTO, s.productClient)
	if err != nil {
		return nil, err
	}

	res := utils.MapOrderResponseDTOToOrderRPC(ordersWithProduct)
	return res, nil
}

func (s *orderService) CreateOrder(ctx context.Context, order dto.CreateOrderRequestDto) (*orderpb.OrderResponse, error) {
	createdOrder, err := s.orderRepository.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	for i := range order.Item {
		order.Item[i].OrderID = createdOrder.ID
	}
	createdItem, err := s.orderItemRepository.CreateOrderItems(ctx, order.Item)
	createdOrder.OrderItems = createdItem

	createdOrder, err = utils.AttachProductDetailToOrder(ctx, createdOrder, s.productClient)
	if err != nil {
		return nil, err
	}

	res := utils.MapOrderResponseDTOToOrderRPC(createdOrder)
	return res, nil
}

func (s *orderService) UpdateStatusOrder(ctx context.Context, order dto.UpdateStatusOrderRequestDto) (*orderpb.OrderResponse, error) {
	updatedOrder, err := s.orderRepository.UpdateStatusOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	updatedOrder, err = utils.AttachProductDetailToOrder(ctx, updatedOrder, s.productClient)
	if err != nil {
		return nil, err
	}

	res := utils.MapOrderResponseDTOToOrderRPC(updatedOrder)
	return res, nil
}

func (s *orderService) DeleteOrder(ctx context.Context, orderId uint64, userId uint64) error {
	return s.orderRepository.DeleteOrder(ctx, orderId, userId)
}
