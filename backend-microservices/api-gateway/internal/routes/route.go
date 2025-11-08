package routes

import (
	userController "api-gateway/internal/controllers/user"
	userpb "api-gateway/internal/pb/user"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userClient userpb.UserServiceClient, addressClient userpb.AddressServiceClient) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api/v1")

	// 🔹 Initializating controller from each service
	userCtrl := userController.NewUserController(userClient)
	addressCtrl := userController.NewAddressController(addressClient)

	// 🔹 Registrasi route
	UserRoute(api, userCtrl, addressCtrl)

	return r
}
