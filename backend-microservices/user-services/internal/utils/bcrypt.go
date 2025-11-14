package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func MatchPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
