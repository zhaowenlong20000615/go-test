package middleware

import (
	"encoding/gob"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go-test/webook/internal/constants"
	"go-test/webook/internal/web"
	"go-test/webook/pkg"
	"net/http"
	"strings"
	"time"
)

type LoginJwtMiddlewareBuilder struct {
	paths []string
}

func NewLoginJwtMiddlewareBuilder() *LoginJwtMiddlewareBuilder {
	return &LoginJwtMiddlewareBuilder{}
}

func (l *LoginJwtMiddlewareBuilder) IgnorePaths(paths ...string) *LoginJwtMiddlewareBuilder {
	for _, path := range paths {
		l.paths = append(l.paths, path)
	}
	return l
}

func (l *LoginJwtMiddlewareBuilder) Build() gin.HandlerFunc {
	gob.Register(time.Now())
	return func(ctx *gin.Context) {
		for _, path := range l.paths {
			if ctx.Request.URL.Path == path {
				return
			}
		}
		tokenStr := ctx.GetHeader("Authorization")
		if tokenStr == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		strArr := strings.Split(tokenStr, " ")
		claims := web.UserClaims{}
		token, err := jwt.ParseWithClaims(strArr[1], &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(constants.SHORT_TIME_JWT_KEY), nil
		})

		isExpired(ctx, err, func() {
			longTokenStr, err2 := pkg.Redis.Client.Get(ctx, token.Raw).Result()
			if err2 != nil || longTokenStr == "" {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			longToken, err := jwt.Parse(longTokenStr, func(token *jwt.Token) (interface{}, error) {
				return []byte(constants.LONG_TIME_JWT_KEY), nil
			})
			isExpired(ctx, err, func() {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"err": err.Error(),
				})
				println("已过期 longToken", longToken)
				return
			})
			newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, web.UserClaims{
				Id: claims.Id,
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
				},
			})
			shortToken, err := newToken.SignedString([]byte(constants.SHORT_TIME_JWT_KEY))
			ctx.Header("x-jwt-token", shortToken)
			ttl, err := pkg.Redis.Client.TTL(ctx, token.Raw).Result()
			pkg.Redis.Client.Del(ctx, token.Raw)
			pkg.Redis.Client.Set(ctx, constants.SHORT_TIME_JWT_KEY, shortToken, time.Hour)
			pkg.Redis.Client.Set(ctx, shortToken, longTokenStr, ttl)
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":  "100",
				"token": shortToken,
			})
		})

		ctx.Set(constants.CLAIMS_KEY, claims)
	}
}

func isExpired(ctx *gin.Context, err error, cb func()) {
	if err == nil {
		return
	}
	if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		if cb != nil {
			cb()
		}
	} else {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"err": err.Error(),
		})
		return
	}
}
