package repositories

import (
	"context"

	"order-services/dto"
	"order-services/models"

	"gorm.io/gorm"
)

type StatusHistoryRepository interface {
	GetStatusHistoryByOrderID(ctx context.Context, orderID uint) ([]dto.StatusHistoryResponseDto, error)
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

func (r *statusHistoryRepository) GetStatusHistoryByOrderID(ctx context.Context, orderID uint) ([]dto.StatusHistoryResponseDto, error) {
	var statusHistory []dto.StatusHistoryResponseDto
	if err := r.db.Where("order_id = ?", orderID).Find(&statusHistory).Error; err != nil {
		return nil, err
	}
	return statusHistory, nil
}

func (r *statusHistoryRepository) CreateStatusHistory(ctx context.Context, statusHistory dto.CreateStatusHistoryRequestDto) (dto.StatusHistoryResponseDto, error) {
	newStatusHistory := models.StatusHistory{
		OrderID: statusHistory.OrderID,
		Status:  statusHistory.Status,
	}
	if err := r.db.Create(&newStatusHistory).Error; err != nil {
		return dto.StatusHistoryResponseDto{}, err
	}
	return dto.StatusHistoryResponseDto{
		Status: newStatusHistory.Status,
		Date:   newStatusHistory.Date.Format("2006-01-02"),
	}, nil
}
