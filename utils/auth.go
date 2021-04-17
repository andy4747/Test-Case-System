package utils

import (
	"os"
	"time"
	"github.com/dgrijalva/jwt-go"
    "fmt"
)

type Credentials struct {
    Email string	`json:"email"`
	Password string	`json:"password"`
}

func GenerateToken(email string) string {
    claims := jwt.MapClaims{
        "exp" : time.Now().Add(time.Hour * 3).Unix(),
        "iat" : time.Now().Unix(),
        "user": email,
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    t, _ := token.SignedString(os.Getenv("SECRET_KEY"))
    return t
}

func ValidateToken(token string) (*jwt.Token, error) {
    return jwt.Parse(token, func (token *jwt.Token) (interface{}, error){
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            //return nil as secret key
            return nil, fmt.Errorf("wring signing method: %v",token.Header["alg"])
        }
        return []byte(os.Getenv("SECRET_KEY")), nil
    })
}

