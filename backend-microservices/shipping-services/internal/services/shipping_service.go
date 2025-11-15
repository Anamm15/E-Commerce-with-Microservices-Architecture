package services

import (
	"context"

	"shipping-service/internal/algorithm"
	"shipping-service/internal/repositories"
	"shipping-service/internal/utils"
	shippingpb "shipping-service/pb"
)

type ShippingService interface {
	GetDetailShipment(ctx context.Context, id uint64) (*shippingpb.ShipmentResponse, error)
	CalculateCostShipping(ctx context.Context, req *shippingpb.CalculateCostShippingRequest) (*shippingpb.CalculateCostShippingListResponse, error)
}

type shippingService struct {
	shippingRepository repositories.ShippingRepository
}

func NewShippingService(shippingRepository repositories.ShippingRepository) ShippingService {
	return &shippingService{shippingRepository: shippingRepository}
}

func (s *shippingService) GetDetailShipment(ctx context.Context, id uint64) (*shippingpb.ShipmentResponse, error) {
	shipment, err := s.shippingRepository.GetDetailShipment(ctx, id)
	if err != nil {
		return nil, err
	}

	shipmentRes := utils.MapShipmentResponseTogRPC(shipment)
	return shipmentRes, nil
}

func (s *shippingService) CalculateCostShipping(ctx context.Context, req *shippingpb.CalculateCostShippingRequest) (*shippingpb.CalculateCostShippingListResponse, error) {
	calculatedCost, err := algorithm.CalculateAllCosts(req.Weight)
	if err != nil {
		return nil, err
	}

	var costList []*shippingpb.CalculateCostShippingResponse
	for _, cost := range calculatedCost {
		costList = append(costList, &shippingpb.CalculateCostShippingResponse{
			Courier:       cost.Courier,
			ServiceType:   cost.ServiceType,
			Price:         float32(cost.TotalCost),
			EstimatedDays: int32(cost.EstimatedDays),
		})
	}

	return &shippingpb.CalculateCostShippingListResponse{Cost: costList}, nil
}
