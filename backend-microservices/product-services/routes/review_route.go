package routes

import (
	"product-services/controllers"

	"github.com/gin-gonic/gin"
)

func ReviewRoutes(server *gin.Engine, reviewController controllers.ReviewController) {
	reviewGroup := server.Group("/api/v1/reviews")
	{
		reviewGroup.GET("/", reviewController.GetAllReviews)
		reviewGroup.GET("/:productId", reviewController.GetReviewByProductID)
		reviewGroup.POST("/", reviewController.CreateReview)
		reviewGroup.PATCH("/:id", reviewController.UpdateReview)
		reviewGroup.DELETE("/:id", reviewController.DeleteReview)
	}
}
