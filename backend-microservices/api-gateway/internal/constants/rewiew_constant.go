package constants

const (
	ContextKeyUserID = "user_id"

	ErrReviewIDOrLoginRequired = "review id or login is required"
	ErrReviewIDRequired        = "review id is required"
	ErrGetReviews              = "Failed to get reviews"
	ErrGetReviewsByProductID   = "Failed to get reviews by product id"
	ErrCreateReview            = "Failed to create review"
	ErrUpdateReview            = "Failed to update review"
	ErrDeleteReview            = "Failed to delete review"

	SuccessGetReviews   = "Reviews retrieved successfully"
	SuccessCreateReview = "Review created successfully"
	SuccessUpdateReview = "Review updated successfully"
	SuccessDeleteReview = "Review deleted successfully"
)
