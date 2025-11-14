package routes

import (
	userController "api-gateway/internal/controllers/user"

	"github.com/gin-gonic/gin"
)

// UserRoute meregistrasi semua endpoint yang berhubungan dengan user
func UserRoute(router *gin.RouterGroup, userController *userController.UserController, addressController *userController.AddressController) {
	user := router.Group("/users")
	{
		user.GET("/", userController.GetAllUsers)
		user.GET("", userController.GetUserByUsername)
		user.POST("/", userController.RegisterUser)
		user.POST("/login", userController.LoginUser)
		user.PATCH("/:id", userController.UpdateUser)
		user.DELETE("/:id", userController.DeleteUser)

		userAddress := user.Group("/address")
		{
			userAddress.GET("/", addressController.GetUserAddress)
			userAddress.POST("/", addressController.CreateUserAddress)
			userAddress.PATCH("/:id", addressController.UpdateUserAddress)
			userAddress.DELETE("/:id", addressController.DeleteUserAddress)
		}
	}
}
