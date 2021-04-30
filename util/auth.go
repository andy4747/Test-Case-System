package util

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	jwt "github.com/dgrijalva/jwt-go"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(userID uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix()
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		log.Errorln("Couldn't generate token")
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			//return nil as secret key
			return nil, fmt.Errorf("wring signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
}
