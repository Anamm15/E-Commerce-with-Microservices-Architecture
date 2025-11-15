package routes

import (
	orderController "api-gateway/internal/controllers/order"
	categoryController "api-gateway/internal/controllers/product"
	productController "api-gateway/internal/controllers/product"
	reviewController "api-gateway/internal/controllers/product"
	userController "api-gateway/internal/controllers/user"
	orderpb "api-gateway/internal/pb/order"
	productpb "api-gateway/internal/pb/product"
	userpb "api-gateway/internal/pb/user"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	userClient userpb.UserServiceClient,
	addressClient userpb.AddressServiceClient,
	categoryClient productpb.CategoryServiceClient,
	productClient productpb.ProductServiceClient,
	reviewClient productpb.ReviewServiceClient,
	orderClient orderpb.OrderServiceClient,
) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api/v1")

	// 🔹 Initializating controller from each service
	userCtrl := userController.NewUserController(userClient)
	addressCtrl := userController.NewAddressController(addressClient)

	// 🔹 Registrasi route
	categoryCtrl := categoryController.NewCategoryController(categoryClient)
	productCtrl := productController.NewProductController(productClient)
	reviewCtrl := reviewController.NewReviewController(reviewClient)
	orderCtrl := orderController.NewOrderController(orderClient)

	// 🔹 Registrasi route
	UserRoute(api, userCtrl, addressCtrl)
	CategoryRoute(api, categoryCtrl)
	ProductRoute(api, productCtrl)
	ReviewRoute(api, reviewCtrl)
	OrderRoute(api, orderCtrl)

	return r
}
