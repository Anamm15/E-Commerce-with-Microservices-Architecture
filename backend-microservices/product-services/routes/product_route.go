package routes

import (
	"product-services/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(server *gin.Engine, productController controllers.ProductController) {
	productGroup := server.Group("/api/v1/products")
	{
		productGroup.GET("/", productController.GetAllProducts)
		productGroup.GET("/:id", productController.GetProductById)
		productGroup.GET("/category/:id", productController.GetProductByCategoryID)
		productGroup.POST("/", productController.CreateProduct)
		productGroup.PATCH("/:id", productController.UpdateProduct)
		productGroup.DELETE("/:id", productController.DeleteProduct)
	}
}
