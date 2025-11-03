package routes

import (
	"product-services/controllers"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(server *gin.Engine, categoryController controllers.CategoryController) {
	categoryGroup := server.Group("/api/v1/categories")
	{
		categoryGroup.GET("/", categoryController.GetAllCategories)
		categoryGroup.POST("/", categoryController.CreateCategory)
		categoryGroup.PATCH("/:id", categoryController.UpdateCategory)
		categoryGroup.DELETE("/:id", categoryController.DeleteCategory)
	}
}
