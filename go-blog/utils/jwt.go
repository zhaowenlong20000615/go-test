package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"go-test/go-blog/models"
	"time"
)

type MyCustomClaims struct {
	User models.User `json:"user"`
	jwt.RegisteredClaims
}

const JWT_KEY = "token"

func CrateToken(user models.User, expireTime time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyCustomClaims{
		User: user,
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

func ParseToken(tokenString string) (*MyCustomClaims, error) {
	myCustomClaims := MyCustomClaims{}
	jwtToken, err := jwt.ParseWithClaims(tokenString, &myCustomClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_KEY), nil
	})
	if err != nil {
		return nil, err
	}
	if jwtToken == nil {
		return nil, nil
	}
	claims, ok := jwtToken.Claims.(*MyCustomClaims)
	if ok && jwtToken.Valid {
		return claims, nil
	}
	return nil, err
}
