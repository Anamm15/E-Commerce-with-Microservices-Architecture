package routes

import (
	"order-services/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(server *gin.Engine, orderController controllers.OrderController) {
	user := server.Group("/api/v1/orders")
	{
		user.GET("/", orderController.GetOrders)
		user.GET("/user/:id", orderController.GetOrdersByUser)
		user.GET("/status/:status", orderController.GetOrdersByStatus)
		user.GET("/:id", orderController.GetDetailOrder)
		user.POST("/", orderController.CreateOrder)
		user.PATCH("/status/:id", orderController.UpdateStatusOrder)
		user.DELETE("/:id", orderController.DeleteOrder)
	}
}
