package utils

import (
	"order-services/internal/dto"
	"order-services/internal/models"
	orderpb "order-services/pb/order"
)

func MapItemRPCToItemDTO(item []*orderpb.CreateOrderItem) []dto.CreateOrderItemRequestDto {
	var res []dto.CreateOrderItemRequestDto
	for _, orderItem := range item {
		res = append(res, dto.CreateOrderItemRequestDto{
			ProductID: orderItem.ProductId,
			Quantity:  orderItem.Quantity,
			Total:     orderItem.Total,
		})
	}

	return res
}

func MapItemDTOToOrderItem(item []dto.OrderItemResponseDto) []*orderpb.OrderItem {
	var res []*orderpb.OrderItem
	for _, orderItem := range item {
		res = append(res, &orderpb.OrderItem{
			Id:        orderItem.ID,
			ProductId: orderItem.ProductID,
			Quantity:  orderItem.Quantity,
			Total:     orderItem.Total,
			Name:      orderItem.Name,
			ImageUrl:  orderItem.ImageUrl,
		})
	}

	return res
}

func MapOrderResponseDTOToOrderListRPC(orders []dto.OrderResponseDto) []*orderpb.OrderResponse {
	var res []*orderpb.OrderResponse
	for _, order := range orders {
		item := MapItemDTOToOrderItem(order.OrderItems)
		res = append(res, &orderpb.OrderResponse{
			Id:                order.ID,
			Status:            order.Status,
			Date:              order.Date,
			ShippingCost:      order.ShippingCost,
			EstimatedDelivery: order.EstimatedDelivery,
			Total:             order.Total,
			PaymentMethod:     order.PaymentMethod,
			TrackingNumber:    order.TrackingNumber,
			OrderItems:        item,
		})
	}

	return res
}

func MapOrderResponseDTOToOrderRPC(orders dto.OrderResponseDto) *orderpb.OrderResponse {
	return &orderpb.OrderResponse{
		Id:                orders.ID,
		Status:            orders.Status,
		Date:              orders.Date,
		ShippingCost:      orders.ShippingCost,
		EstimatedDelivery: orders.EstimatedDelivery,
		Total:             orders.Total,
		PaymentMethod:     orders.PaymentMethod,
		TrackingNumber:    orders.TrackingNumber,
		OrderItems:        MapItemDTOToOrderItem(orders.OrderItems),
	}
}

func MapOrderModelToDTO(order models.Order) dto.OrderResponseDto {
	itemDTOs := make([]dto.OrderItemResponseDto, len(order.OrderItems))
	for i, it := range order.OrderItems {
		itemDTOs[i] = dto.OrderItemResponseDto{
			ProductID: it.ProductID,
			Quantity:  it.Quantity,
			Total:     it.Total,
		}
	}

	return dto.OrderResponseDto{
		ID:                order.ID,
		Status:            order.Status,
		ShippingCost:      order.ShippingCost,
		EstimatedDelivery: order.EstimatedDelivery,
		Total:             order.Total,
		PaymentMethod:     order.PaymentMethod,
		TrackingNumber:    order.TrackingNumber,
		OrderItems:        itemDTOs,
	}
}

func MapOrdersModelToDTO(orders []models.Order) []dto.OrderResponseDto {
	result := make([]dto.OrderResponseDto, len(orders))

	for i, order := range orders {
		result[i] = MapOrderModelToDTO(order)
	}

	return result
}
