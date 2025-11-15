package repositories

import (
	"context"
	"errors"

	"order-services/internal/constants"
	"order-services/internal/dto"
	"order-services/internal/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	GetOrders(ctx context.Context) ([]models.Order, error)
	GetOrdersByUser(ctx context.Context, userID uint64) ([]models.Order, error)
	GetOrdersByStatus(ctx context.Context, status string) ([]models.Order, error)
	GetDetailOrder(ctx context.Context, orderID uint64) (models.Order, error)
	CreateOrder(ctx context.Context, order dto.CreateOrderRequestDto) (dto.OrderResponseDto, error)
	UpdateStatusOrder(ctx context.Context, order dto.UpdateStatusOrderRequestDto) (dto.OrderResponseDto, error)
	DeleteOrder(ctx context.Context, orderId uint64, userId uint64) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) GetOrders(ctx context.Context) ([]models.Order, error) {
	var orders []models.Order
	if err := r.db.WithContext(ctx).
		Preload("OrderItems").
		Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *orderRepository) GetOrdersByUser(ctx context.Context, userID uint64) ([]models.Order, error) {
	var orders []models.Order
	if err := r.db.WithContext(ctx).
		Preload("OrderItems").
		Where("user_id = ?", userID).
		Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *orderRepository) GetOrdersByStatus(ctx context.Context, status string) ([]models.Order, error) {
	var orders []models.Order
	if err := r.db.WithContext(ctx).
		Preload("OrderItems").
		Where("status = ?", status).
		Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *orderRepository) GetDetailOrder(ctx context.Context, orderID uint64) (models.Order, error) {
	var order models.Order
	if err := r.db.WithContext(ctx).
		Model(&models.Order{}).
		Preload("OrderItems").
		Where("id = ?", orderID).
		Take(&order).Error; err != nil {
		return models.Order{}, err
	}

	return order, nil
}

func (r *orderRepository) CreateOrder(ctx context.Context, order dto.CreateOrderRequestDto) (dto.OrderResponseDto, error) {
	newOrder := models.Order{
		UserID:        order.UserID,
		Total:         order.Total,
		Status:        constants.ORDER_STATUS_PENDING,
		PaymentMethod: order.PaymentMethod,
	}

	if err := r.db.WithContext(ctx).
		Create(&newOrder).Error; err != nil {
		return dto.OrderResponseDto{}, err
	}

	return dto.OrderResponseDto{
		ID:     newOrder.ID,
		Status: newOrder.Status,
	}, nil
}

func (r *orderRepository) UpdateStatusOrder(ctx context.Context, req dto.UpdateStatusOrderRequestDto) (dto.OrderResponseDto, error) {
	result := r.db.WithContext(ctx).
		Model(&models.Order{}).
		Where("id = ?", req.OrderID).
		Updates(map[string]interface{}{
			"status": req.Status,
		})

	if result.Error != nil {
		return dto.OrderResponseDto{}, result.Error
	}

	if result.RowsAffected == 0 {
		return dto.OrderResponseDto{}, errors.New("order not found")
	}

	return dto.OrderResponseDto{
		ID:     req.OrderID,
		Status: req.Status,
	}, nil
}

func (r *orderRepository) DeleteOrder(ctx context.Context, orderId uint64, userId uint64) error {
	result := r.db.WithContext(ctx).
		Where("id = ? AND user_id = ?", orderId, userId).
		Delete(&models.Order{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("order not found or unauthorized")
	}

	return nil
}
