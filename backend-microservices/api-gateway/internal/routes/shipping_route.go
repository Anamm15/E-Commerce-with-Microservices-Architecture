package routes

import (
	shippingController "api-gateway/internal/controllers/shipping"

	"github.com/gin-gonic/gin"
)

func ShippingRoute(router *gin.RouterGroup, shippingController *shippingController.ShippingController) {
	shipping := router.Group("/shipments")
	{
		shipping.GET("/:id", shippingController.GetDetailShippment)
		shipping.POST("/cost", shippingController.CalculateCostShipping)
	}
}
