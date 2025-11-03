package repositories

import (
	"context"

	"order-services/dto"
	"order-services/models"

	"gorm.io/gorm"
)

type OrderItemRepository interface {
	getOrderItemsByOrderID(ctx context.Context, orderID uint) ([]dto.OrderItemResponseDto, error)
	createOrderItem(ctx context.Context, orderItem dto.CreateOrderItemRequestDto) (dto.OrderItemResponseDto, error)
}

type orderItemRepository struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) OrderItemRepository {
	return &orderItemRepository{
		db: db,
	}
}

func (r *orderItemRepository) getOrderItemsByOrderID(ctx context.Context, orderID uint) ([]dto.OrderItemResponseDto, error) {
	var orderItems []dto.OrderItemResponseDto
	if err := r.db.Where("order_id = ?", orderID).Find(&orderItems).Error; err != nil {
		return nil, err
	}
	return orderItems, nil
}

func (r *orderItemRepository) createOrderItem(ctx context.Context, orderItem dto.CreateOrderItemRequestDto) (dto.OrderItemResponseDto, error) {
	newOrderItem := models.OrderItem{
		OrderID:   orderItem.OrderID,
		ProductID: orderItem.ProductID,
		Quantity:  orderItem.Quantity,
		Total:     orderItem.Total,
	}
	if err := r.db.Create(&newOrderItem).Error; err != nil {
		return dto.OrderItemResponseDto{}, err
	}
	return dto.OrderItemResponseDto{
		ID:        newOrderItem.ID,
		OrderID:   newOrderItem.OrderID,
		ProductID: newOrderItem.ProductID,
		Quantity:  newOrderItem.Quantity,
		Total:     newOrderItem.Total,
	}, nil
}
