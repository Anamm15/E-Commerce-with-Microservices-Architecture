package routes

import (
	categoryController "api-gateway/internal/controllers/product"
	productController "api-gateway/internal/controllers/product"
	reviewController "api-gateway/internal/controllers/product"
	userController "api-gateway/internal/controllers/user"
	productpb "api-gateway/internal/pb/product"
	userpb "api-gateway/internal/pb/user"

	"github.com/gin-gonic/gin"
)

// func SetupRouter(authClient userpb.AuthServiceClient, userClient userpb.UserServiceClient) *gin.Engine {
func SetupRouter(
	userClient userpb.UserServiceClient,
	categoryClient productpb.CategoryServiceClient,
	productClient productpb.ProductServiceClient,
	reviewClient productpb.ReviewServiceClient,
) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api/v1")

	// 🔹 Inisialisasi controller dari subfolder
	userCtrl := userController.NewUserController(userClient)
	categoryCtrl := categoryController.NewCategoryController(categoryClient)
	productCtrl := productController.NewProductController(productClient)
	reviewCtrl := reviewController.NewReviewController(reviewClient)

	// 🔹 Registrasi route
	UserRoute(api, userCtrl)
	CategoryRoute(api, categoryCtrl)
	ProductRoute(api, productCtrl)
	ReviewRoute(api, reviewCtrl)

	return r
}
