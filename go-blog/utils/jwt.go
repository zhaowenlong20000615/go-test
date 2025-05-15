package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type MyCustomClaims struct {
	user interface{}
	jwt.RegisteredClaims
}

const JWT_KEY = "token"

func CrateToken(user interface{}, expireTime time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyCustomClaims{
		user: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireTime)),
		},
	})
	signedString, err := token.SignedString([]byte(JWT_KEY))
	if err != nil {
		return "", err
	}
	return signedString, nil
}

func ParseToken() {

}
