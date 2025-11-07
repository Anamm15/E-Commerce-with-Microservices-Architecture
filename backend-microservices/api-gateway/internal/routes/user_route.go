package routes

import (
	userController "api-gateway/internal/controllers/user"

	"github.com/gin-gonic/gin"
)

// UserRoute meregistrasi semua endpoint yang berhubungan dengan user
func UserRoute(router *gin.RouterGroup, userController *userController.UserController) {
	user := router.Group("/users")
	{
		user.GET("/", userController.GetAllUsers)
		user.GET("/:username", userController.GetUserByUsername)
		user.POST("/", userController.CreateUser)
		user.PATCH("/:id", userController.UpdateUser)
		user.DELETE("/:id", userController.DeleteUser)
	}
}
