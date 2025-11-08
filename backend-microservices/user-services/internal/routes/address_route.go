package routes

import (
	"user-services/internal/controllers"

	"github.com/gin-gonic/gin"
)

func AddressRoute(server *gin.Engine, addressController controllers.AddressController) {
	user := server.Group("/api/v1/user/address")
	{
		user.GET("/", addressController.GetUserAddress)
		user.POST("/", addressController.CreateUserAddress)
		user.PUT("/:id", addressController.UpdateUserAddress)
		user.DELETE("/:id", addressController.DeleteUserAddress)
	}
}
