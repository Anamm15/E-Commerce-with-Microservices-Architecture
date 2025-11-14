package utils

import (
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key")

func ParseTokenJWT(tokenString string) (uint64, string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// pastikan metode signing benar
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})
	if err != nil {
		return 0, "", err
	}

	// Ambil claims (MapClaims)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, "", jwt.ErrTokenInvalidClaims
	}

	// Extract user_id
	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, "", jwt.ErrTokenInvalidClaims
	}
	userID := uint64(userIDFloat)

	// Extract role
	role, ok := claims["role"].(string)
	if !ok {
		return 0, "", jwt.ErrTokenInvalidClaims
	}

	return userID, role, nil
}
