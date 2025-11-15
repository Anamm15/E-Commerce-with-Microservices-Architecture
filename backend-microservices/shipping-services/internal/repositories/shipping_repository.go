package repositories

import (
	"context"

	"shipping-service/internal/models"

	"gorm.io/gorm"
)

type ShippingRepository interface {
	GetDetailShipment(ctx context.Context, id uint64) (models.Shipment, error)
	CreateShipment(ctx context.Context, shipment models.Shipment) (models.Shipment, error)
	UpdateShipment(ctx context.Context, id uint64, shipment models.Shipment) (models.Shipment, error)
}

type shippingRepository struct {
	db *gorm.DB
}

func NewShippingRepository(db *gorm.DB) ShippingRepository {
	return &shippingRepository{db: db}
}

func (r *shippingRepository) GetDetailShipment(ctx context.Context, id uint64) (models.Shipment, error) {
	var shipment models.Shipment
	if err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&shipment).Error; err != nil {
		return models.Shipment{}, err
	}
	return shipment, nil
}

func (r *shippingRepository) CreateShipment(ctx context.Context, shipment models.Shipment) (models.Shipment, error) {
	if err := r.db.WithContext(ctx).
		Create(&shipment).Error; err != nil {
		return models.Shipment{}, err
	}
	return shipment, nil
}

func (r *shippingRepository) UpdateShipment(ctx context.Context, id uint64, shipment models.Shipment) (models.Shipment, error) {
	if err := r.db.WithContext(ctx).
		Where("id = ?", id).
		Updates(&shipment).Error; err != nil {
		return models.Shipment{}, err
	}
	return shipment, nil
}
