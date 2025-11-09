package routes

import (
	controllers "api-gateway/internal/controllers/product"

	"github.com/gin-gonic/gin"
)

// UserRoute meregistrasi semua endpoint yang berhubungan dengan user
func CategoryRoute(router *gin.RouterGroup, categoryController controllers.CategoryController) {
	categoryGroup := router.Group("/categories")
	{
		categoryGroup.GET("/", categoryController.GetAllCategories)
		categoryGroup.POST("/", categoryController.CreateCategory)
		categoryGroup.PATCH("/:id", categoryController.UpdateCategory)
		categoryGroup.DELETE("/:id", categoryController.DeleteCategory)
	}
}

func ProductRoute(router *gin.RouterGroup, productController controllers.ProductController) {
	productGroup := router.Group("/products")
	{
		productGroup.GET("/", productController.GetAllProducts)
		productGroup.GET("/:id", productController.GetProductById)
		productGroup.GET("/category/:id", productController.GetProductByCategoryID)
		productGroup.POST("/", productController.CreateProduct)
		productGroup.PATCH("/:id", productController.UpdateProduct)
		productGroup.DELETE("/:id", productController.DeleteProduct)
	}
}

func ReviewRoute(router *gin.RouterGroup, reviewController controllers.ReviewController) {
	reviewGroup := router.Group("/api/v1/reviews")
	{
		reviewGroup.GET("/", reviewController.GetAllReviews)
		reviewGroup.GET("/:productId", reviewController.GetReviewByProductID)
		reviewGroup.POST("/", reviewController.CreateReview)
		reviewGroup.PATCH("/:id", reviewController.UpdateReview)
		reviewGroup.DELETE("/:id", reviewController.DeleteReview)
	}
}
