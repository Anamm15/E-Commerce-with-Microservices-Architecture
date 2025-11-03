package constants

import (
	"errors"
)

const (
	ADDRESS_RETRIEVED_SUCCESSFULLY = "Address retrieved successfully."
	ADDRESS_CREATED_SUCCESSFULLY   = "Address created successfully."
	ADDRESS_UPDATED_SUCCESSFULLY   = "Address updated successfully."
	ADDRESS_DELETED_SUCCESSFULLY   = "Address deleted successfully."
	ADDRESS_NOT_FOUND              = "Address not found."
	ADDRESS_CREATION_FAILED        = "Failed to create address."
	ADDRESS_UPDATE_FAILED          = "Failed to update address."
	ADDRESS_DELETION_FAILED        = "Failed to delete address."
)

var (
	ErrAddressNotFound       = errors.New("address not found")
	ErrFailedToCreateAddress = errors.New("failed to create address")
	ErrFailedToUpdateAddress = errors.New("failed to update address")
	ErrFailedToDeleteAddress = errors.New("failed to delete address")
)
