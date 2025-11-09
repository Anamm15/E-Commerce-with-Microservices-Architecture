package routes

import (
	addressController "api-gateway/internal/controllers/user"

	"github.com/gin-gonic/gin"
)

func AddressRoute(server *gin.RouterGroup, addressController *addressController.AddressController) {
	user := server.Group("/user/address")
	{
		user.GET("/", addressController.GetUserAddress)
		user.POST("/", addressController.CreateUserAddress)
		user.PUT("/:id", addressController.UpdateUserAddress)
		user.DELETE("/:id", addressController.DeleteUserAddress)
	}
}
