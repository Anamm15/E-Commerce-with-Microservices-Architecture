package constants

import (
	"errors"
)

const (
	USER_RETRIEVED_SUCCESSFULLY = "User retrieved successfully."
	USER_CREATED_SUCCESSFULLY   = "User created successfully."
	USER_UPDATED_SUCCESSFULLY   = "User updated successfully."
	USER_DELETED_SUCCESSFULLY   = "User deleted successfully."
	USER_NOT_FOUND              = "User not found."
	USER_CREATION_FAILED        = "Failed to create user."
	USER_UPDATE_FAILED          = "Failed to update user."
	USER_DELETION_FAILED        = "Failed to delete user."
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrFailedToCreateUser = errors.New("failed to create user")
	ErrFailedToUpdateUser = errors.New("failed to update user")
	ErrFailedToDeleteUser = errors.New("failed to delete user")
)
