package controllers

import (
	"context"

	"shipping-service/internal/constants"
	"shipping-service/internal/services"
	"shipping-service/pb"
	shippingpb "shipping-service/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ShippingController struct {
	shippingpb.UnimplementedShippingServiceServer
	shippingService services.ShippingService
}

func NewShippingController(shippingService services.ShippingService) *ShippingController {
	return &ShippingController{shippingService: shippingService}
}

func (c *ShippingController) GetDetailShipment(ctx context.Context, req *shippingpb.GetDetailShipmentRequest) (*pb.ShipmentResponse, error) {
	shipment, err := c.shippingService.GetDetailShipment(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, constants.ErrDetailShipmentGet, err)
	}
	return shipment, nil
}

func (c *ShippingController) CalculateCostShipping(ctx context.Context, req *shippingpb.CalculateCostShippingRequest) (*pb.CalculateCostShippingListResponse, error) {
	cost, err := c.shippingService.CalculateCostShipping(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, constants.ErrCalculateCostShipping, err)
	}
	return cost, nil
}
