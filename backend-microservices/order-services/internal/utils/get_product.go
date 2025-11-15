package utils

import (
	"context"

	"order-services/internal/dto"

	productpb "order-services/pb/product"
)

func attachProductDetail(
	ctx context.Context,
	order *dto.OrderResponseDto,
	productClient productpb.ProductServiceClient,
) error {
	for i, item := range order.OrderItems {
		product, err := productClient.GetProductByID(ctx, &productpb.GetProductByIDRequest{
			Id: item.ProductID,
		})
		if err != nil {
			return err
		}

		order.OrderItems[i].Name = product.Name

		if len(product.ImageUrl) > 0 {
			order.OrderItems[i].ImageUrl = product.ImageUrl[0].Url
		}
	}

	return nil
}

func AttachProductDetailToOrders(
	ctx context.Context,
	orders []dto.OrderResponseDto,
	productClient productpb.ProductServiceClient,
) ([]dto.OrderResponseDto, error) {
	for i := range orders {
		if err := attachProductDetail(ctx, &orders[i], productClient); err != nil {
			return nil, err
		}
	}

	return orders, nil
}

func AttachProductDetailToOrder(
	ctx context.Context,
	order dto.OrderResponseDto,
	productClient productpb.ProductServiceClient,
) (dto.OrderResponseDto, error) {
	if err := attachProductDetail(ctx, &order, productClient); err != nil {
		return dto.OrderResponseDto{}, err
	}

	return order, nil
}
