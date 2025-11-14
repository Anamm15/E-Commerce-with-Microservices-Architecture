package constants

const (
	ErrInvalidRequest     = "invalid request"
	ErrGRPC               = "gRPC error"
	ErrIDRequired         = "id is required"
	ErrInvalidID          = "invalid id"
	ErrUnauthorized       = "unauthorized"
	ErrLoginRequired      = "login is required"
	ErrInvalidToken       = "invalid token"
	ErrRoleNotFound       = "role not found"
	ErrPermissionRequired = "you have no permission to access this resource"
	ErrForbidden          = "forbidden"

	ParamID          = "id"
	QueryUsername    = "username"
	ContextKeyUserID = "user_id"
	ParamAddressID   = "address_id"
)
