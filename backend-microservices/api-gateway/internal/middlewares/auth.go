package middlewares

import (
	"net/http"
	"slices"
	"strings"

	"api-gateway/internal/constants"
	"api-gateway/internal/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, utils.BuildResponseFailed(constants.ErrUnauthorized, constants.ErrLoginRequired, nil))
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, utils.BuildResponseFailed(constants.ErrUnauthorized, constants.ErrInvalidToken, nil))
			c.Abort()
			return
		}

		token := parts[1]

		userID, role, err := utils.ParseTokenJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.BuildResponseFailed(constants.ErrUnauthorized, constants.ErrInvalidToken, nil))
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Set("role", role)
		c.Next()
	}
}

func AuthorizeRoles(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleVal, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, utils.BuildResponseFailed(constants.ErrForbidden, constants.ErrRoleNotFound, nil))
			c.Abort()
			return
		}

		role := roleVal.(string)

		if slices.Contains(allowedRoles, role) {
			c.Next()
			return
		}

		c.JSON(http.StatusForbidden, utils.BuildResponseFailed(constants.ErrForbidden, constants.ErrPermissionRequired, nil))
		c.Abort()
	}
}
