package utils

import (
	"shipping-service/internal/models"

	shippingpb "shipping-service/pb"
)

func MapAddressResponseTogRPC(address models.Address) *shippingpb.Address {
	var addressResponse *shippingpb.Address

	addressResponse = &shippingpb.Address{
		Id:            address.ID,
		RecipientName: address.RecipientName,
		Street:        address.Street,
		Phone:         address.Phone,
		PostalCode:    address.PostalCode,
	}
	return addressResponse
}

func MapShipmentResponseTogRPC(shipment models.Shipment) *shippingpb.ShipmentResponse {
	var shipmentResponse *shippingpb.ShipmentResponse

	shipmentResponse = &shippingpb.ShipmentResponse{
		Id:              shipment.ID,
		Courier:         shipment.Courier,
		ServiceType:     shipment.ServiceType,
		DestinationCity: shipment.DestinationCity,
		Weight:          int32(shipment.Weight),
		Price:           float32(shipment.Price),
		Status:          shipment.Status,
		TrackingNumber:  shipment.TrackingNumber,
		EstimatedDays:   int32(shipment.EstimatedDays),
		Address:         MapAddressResponseTogRPC(shipment.Address),
	}
	return shipmentResponse
}
