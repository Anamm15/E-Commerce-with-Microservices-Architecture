package services

import (
	"context"

	"order-services/dto"
	"order-services/repositories"
)

type OrderService interface {
	GetOrders(ctx context.Context) ([]dto.OrderResponseDto, error)
	GetOrdersByUser(ctx context.Context, userID uint) ([]dto.OrderResponseDto, error)
	GetOrdersByStatus(ctx context.Context, status string) ([]dto.OrderResponseDto, error)
	GetDetailOrder(ctx context.Context, orderID uint) (dto.OrderResponseDto, error)
	CreateOrder(ctx context.Context, order dto.CreateOrderRequestDto) (dto.OrderResponseDto, error)
	UpdateStatusOrder(ctx context.Context, order dto.UpdateStatusOrderRequestDto) (dto.OrderResponseDto, error)
	DeleteOrder(ctx context.Context, orderId uint) error
}

type orderService struct {
	orderRepository repositories.OrderRepository
}

func NewOrderService(orderRepository repositories.OrderRepository) OrderService {
	return &orderService{orderRepository: orderRepository}
}

func (s *orderService) GetOrders(ctx context.Context) ([]dto.OrderResponseDto, error) {
	orders, err := s.orderRepository.GetOrders(ctx)
	return orders, err
}

func (s *orderService) GetOrdersByUser(ctx context.Context, userID uint) ([]dto.OrderResponseDto, error) {
	return s.orderRepository.GetOrdersByUser(ctx, userID)
}

func (s *orderService) GetOrdersByStatus(ctx context.Context, status string) ([]dto.OrderResponseDto, error) {
	return s.orderRepository.GetOrdersByStatus(ctx, status)
}

func (s *orderService) GetDetailOrder(ctx context.Context, orderID uint) (dto.OrderResponseDto, error) {
	return s.orderRepository.GetDetailOrder(ctx, orderID)
}

func (s *orderService) CreateOrder(ctx context.Context, order dto.CreateOrderRequestDto) (dto.OrderResponseDto, error) {
	return s.orderRepository.CreateOrder(ctx, order)
}

func (s *orderService) UpdateStatusOrder(ctx context.Context, order dto.UpdateStatusOrderRequestDto) (dto.OrderResponseDto, error) {
	return s.orderRepository.UpdateStatusOrder(ctx, order)
}

func (s *orderService) DeleteOrder(ctx context.Context, orderId uint) error {
	return s.orderRepository.DeleteOrder(ctx, orderId)
}
