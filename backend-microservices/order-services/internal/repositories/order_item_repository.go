package repositories

import (
	"context"

	"order-services/internal/dto"
	"order-services/internal/models"

	"gorm.io/gorm"
)

type OrderItemRepository interface {
	GetOrderItemsByOrderID(ctx context.Context, orderID uint64) ([]dto.OrderItemResponseDto, error)
	CreateOrderItems(ctx context.Context, itemsReq []dto.CreateOrderItemRequestDto) ([]dto.OrderItemResponseDto, error)
}

type orderItemRepository struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) OrderItemRepository {
	return &orderItemRepository{
		db: db,
	}
}

func (r *orderItemRepository) GetOrderItemsByOrderID(ctx context.Context, orderID uint64) ([]dto.OrderItemResponseDto, error) {
	var orderItems []dto.OrderItemResponseDto
	if err := r.db.WithContext(ctx).
		Model(models.OrderItem{}).
		Where("order_id = ?", orderID).
		Find(&orderItems).Error; err != nil {
		return nil, err
	}
	return orderItems, nil
}

func (r *orderItemRepository) CreateOrderItems(
	ctx context.Context,
	itemsReq []dto.CreateOrderItemRequestDto,
) ([]dto.OrderItemResponseDto, error) {
	orderItems := make([]models.OrderItem, len(itemsReq))

	for i, item := range itemsReq {
		orderItems[i] = models.OrderItem{
			OrderID:   item.OrderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Total:     item.Total,
		}
	}

	if err := r.db.WithContext(ctx).Create(&orderItems).Error; err != nil {
		return nil, err
	}

	responses := make([]dto.OrderItemResponseDto, len(orderItems))
	for i, item := range orderItems {
		responses[i] = dto.OrderItemResponseDto{
			ID:        item.ID,
			OrderID:   item.OrderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Total:     item.Total,
		}
	}

	return responses, nil
}
