package routes

import (
	categoryController "api-gateway/internal/controllers/product"
	productController "api-gateway/internal/controllers/product"
	reviewController "api-gateway/internal/controllers/product"
	shippingController "api-gateway/internal/controllers/shipping"
	userController "api-gateway/internal/controllers/user"
	productpb "api-gateway/internal/pb/product"
	shippingpb "api-gateway/internal/pb/shipping"
	userpb "api-gateway/internal/pb/user"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	userClient userpb.UserServiceClient,
	addressClient userpb.AddressServiceClient,
	categoryClient productpb.CategoryServiceClient,
	productClient productpb.ProductServiceClient,
	reviewClient productpb.ReviewServiceClient,
	shippingClient shippingpb.ShippingServiceClient,
) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api/v1")

	// 🔹 Registrasi route
	userCtrl := userController.NewUserController(userClient)
	addressCtrl := userController.NewAddressController(addressClient)
	categoryCtrl := categoryController.NewCategoryController(categoryClient)
	productCtrl := productController.NewProductController(productClient)
	reviewCtrl := reviewController.NewReviewController(reviewClient)
	shippingCtrl := shippingController.NewShippingController(shippingClient)

	// 🔹 Registrasi route
	UserRoute(api, userCtrl, addressCtrl)
	CategoryRoute(api, categoryCtrl)
	ProductRoute(api, productCtrl)
	ReviewRoute(api, reviewCtrl)
	ShippingRoute(api, shippingCtrl)

	return r
}
