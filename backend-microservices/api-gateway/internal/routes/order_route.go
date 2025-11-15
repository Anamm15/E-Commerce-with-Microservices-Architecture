package routes

import (
	controller "api-gateway/internal/controllers/order"
	"api-gateway/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func OrderRoute(server *gin.RouterGroup, orderController controller.OrderController) {
	user := server.Group("/orders")
	{
		user.GET("/", orderController.GetOrders)
		user.GET("/user/:id", orderController.GetOrdersByUser)
		user.GET("/status/:status", orderController.GetOrdersByStatus)
		user.GET("/:id", orderController.GetDetailOrder)
		user.POST("/", middlewares.AuthMiddleware(), orderController.CreateOrder)
		user.PATCH("/status/:id", orderController.UpdateStatusOrder)
		user.DELETE("/:id", orderController.DeleteOrder)
	}
}
