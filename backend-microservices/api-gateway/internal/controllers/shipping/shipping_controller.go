package shipping

import (
	"net/http"

	"api-gateway/internal/constants"
	dto "api-gateway/internal/dto/shipping"
	shippingpb "api-gateway/internal/pb/shipping"
	"api-gateway/internal/utils"

	"github.com/gin-gonic/gin"
)

type ShippingController struct {
	ShippingClient shippingpb.ShippingServiceClient
}

func NewShippingController(ShippingClient shippingpb.ShippingServiceClient) *ShippingController {
	return &ShippingController{ShippingClient: ShippingClient}
}

func (c *ShippingController) GetDetailShippment(ctx *gin.Context) {
	idParam := ctx.Param("id")
	if idParam == "" {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(constants.ErrInvalidRequest, constants.ErrShippingIDRequired, nil))
		return
	}

	id := utils.StringToUint(idParam)

	grpcReq := &shippingpb.GetDetailShipmentRequest{
		Id: id,
	}

	res, err := c.ShippingClient.GetDetailShipment(ctx, grpcReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(constants.ErrGetDetailShipment, err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess(constants.SuccessGetDetailShipment, res))
}

func (c *ShippingController) CalculateCostShipping(ctx *gin.Context) {
	var req dto.CalculateCostRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(constants.ErrInvalidRequest, err.Error(), nil))
		return
	}

	grpcReq := &shippingpb.CalculateCostShippingRequest{
		OriginCity:      req.OriginCity,
		DestinationCity: req.DestinationCity,
		PostalCode:      req.PostalCode,
		Weight:          req.Weight,
	}

	res, err := c.ShippingClient.CalculateCostShipping(ctx, grpcReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildResponseFailed(constants.ErrCalculateCost, err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponseSuccess(constants.SuccessCalculateCost, res))
}
