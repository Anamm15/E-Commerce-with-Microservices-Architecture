package routes

import (
	userController "api-gateway/internal/controllers/user"
	userpb "api-gateway/internal/pb/user"

	"github.com/gin-gonic/gin"
)

// func SetupRouter(authClient userpb.AuthServiceClient, userClient userpb.UserServiceClient) *gin.Engine {
func SetupRouter(userClient userpb.UserServiceClient) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api/v1")

	// 🔹 Inisialisasi controller dari subfolder
	userCtrl := userController.NewUserController(userClient)

	// 🔹 Registrasi route
	UserRoute(api, userCtrl)

	return r
}
