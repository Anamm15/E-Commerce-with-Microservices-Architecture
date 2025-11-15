package constants

const (
	OrderStatusPending    = "pending"
	OrderStatusProcessing = "processing"
	OrderStatusShipped    = "shipped"
	OrderStatusDelivered  = "delivered"
	OrderStatusCancelled  = "cancelled"

	ORDER_RETRIEVED_SUCCESSFULLY = "Order retrieved successfully"
	ORDER_CREATED_SUCCESSFULLY   = "Order created successfully"
	ORDER_UPDATED_SUCCESSFULLY   = "Order updated successfully"
	ORDER_DELETED_SUCCESSFULLY   = "Order deleted successfully"

	ErrOrderGet      = "Failed to get order %v"
	ErrOrderNotFound = "Order not found %v"
	ErrOrderCreate   = "Failed to create order %v"
	ErrOrderUpdate   = "Failed to update order %v"
	ErrOrderDelete   = "Failed to delete order %v"

	ORDER_STATUS_PENDING = "pending"
)
