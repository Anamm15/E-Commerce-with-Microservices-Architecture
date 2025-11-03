package repositories

import (
	"context"

	"order-services/dto"
	"order-services/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	GetOrders(ctx context.Context) ([]dto.OrderResponseDto, error)
	GetOrdersByUser(ctx context.Context, userID uint) ([]dto.OrderResponseDto, error)
	GetOrdersByStatus(ctx context.Context, status string) ([]dto.OrderResponseDto, error)
	GetDetailOrder(ctx context.Context, orderID uint) (dto.OrderResponseDto, error)
	CreateOrder(ctx context.Context, order dto.CreateOrderRequestDto) (dto.OrderResponseDto, error)
	UpdateStatusOrder(ctx context.Context, order dto.UpdateStatusOrderRequestDto) (dto.OrderResponseDto, error)
	DeleteOrder(ctx context.Context, orderId uint) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) GetOrders(ctx context.Context) ([]dto.OrderResponseDto, error) {
	var orders []dto.OrderResponseDto
	if err := r.db.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *orderRepository) GetOrdersByUser(ctx context.Context, userID uint) ([]dto.OrderResponseDto, error) {
	var orders []dto.OrderResponseDto
	if err := r.db.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *orderRepository) GetOrdersByStatus(ctx context.Context, status string) ([]dto.OrderResponseDto, error) {
	var orders []dto.OrderResponseDto
	if err := r.db.Where("status = ?", status).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *orderRepository) GetDetailOrder(ctx context.Context, orderID uint) (dto.OrderResponseDto, error) {
	var order dto.OrderResponseDto
	if err := r.db.Where("id = ?", orderID).First(&order).Error; err != nil {
		return dto.OrderResponseDto{}, err
	}
	return order, nil
}

func (r *orderRepository) CreateOrder(ctx context.Context, order dto.CreateOrderRequestDto) (dto.OrderResponseDto, error) {
	newOrder := models.Order{
		UserID:            order.UserID,
		Total:             order.Total,
		Status:            "pending",
		ShippingCost:      order.ShippingCost,
		TrackingNumber:    order.TrackingNumber,
		EstimatedDelivery: order.EstimatedDelivery,
		PaymentMethod:     order.PaymentMethod,
	}
	if err := r.db.Create(&newOrder).Error; err != nil {
		return dto.OrderResponseDto{}, err
	}
	return dto.OrderResponseDto{
		ID:     newOrder.ID,
		Status: newOrder.Status,
	}, nil
}

func (r *orderRepository) UpdateStatusOrder(ctx context.Context, order dto.UpdateStatusOrderRequestDto) (dto.OrderResponseDto, error) {
	orderInput := models.Order{
		Status: order.Status,
	}
	if err := r.db.Save(orderInput).Error; err != nil {
		return dto.OrderResponseDto{}, err
	}
	return dto.OrderResponseDto{
		ID:     order.OrderID,
		Status: order.Status,
	}, nil
}

func (r *orderRepository) DeleteOrder(ctx context.Context, orderId uint) error {
	if err := r.db.Delete(&models.Order{}, orderId).Error; err != nil {
		return err
	}
	return nil
}
