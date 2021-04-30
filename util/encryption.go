package util

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password *string) {
	bytePass := []byte(*password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	*password = string(hashedPassword)
}

func ComparePassword(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
