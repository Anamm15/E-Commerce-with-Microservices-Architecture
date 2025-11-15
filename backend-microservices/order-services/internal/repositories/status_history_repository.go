package repositories

import (
	"context"

	"order-services/internal/dto"
	"order-services/internal/models"

	"gorm.io/gorm"
)

type StatusHistoryRepository interface {
	GetStatusHistoryByOrderID(ctx context.Context, orderID uint64) ([]dto.StatusHistoryResponseDto, error)
	CreateStatusHistory(ctx context.Context, statusHistory dto.CreateStatusHistoryRequestDto) (dto.StatusHistoryResponseDto, error)
}

type statusHistoryRepository struct {
	db *gorm.DB
}

func NewStatusHistoryRepository(db *gorm.DB) StatusHistoryRepository {
	return &statusHistoryRepository{
		db: db,
	}
}

func (r *statusHistoryRepository) GetStatusHistoryByOrderID(ctx context.Context, orderID uint64) ([]dto.StatusHistoryResponseDto, error) {
	var statusHistory []dto.StatusHistoryResponseDto
	if err := r.db.WithContext(ctx).
		Model(models.StatusHistory{}).
		Where("order_id = ?", orderID).
		Find(&statusHistory).Error; err != nil {
		return nil, err
	}
	return statusHistory, nil
}

func (r *statusHistoryRepository) CreateStatusHistory(ctx context.Context, statusHistory dto.CreateStatusHistoryRequestDto) (dto.StatusHistoryResponseDto, error) {
	newStatusHistory := models.StatusHistory{
		OrderID: statusHistory.OrderID,
		Status:  statusHistory.Status,
	}
	if err := r.db.WithContext(ctx).Create(&newStatusHistory).Error; err != nil {
		return dto.StatusHistoryResponseDto{}, err
	}
	return dto.StatusHistoryResponseDto{
		Status: newStatusHistory.Status,
		Date:   newStatusHistory.Date.Format("2006-01-02"),
	}, nil
}
