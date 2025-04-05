package middleware

import (
	"encoding/gob"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-test/webook/internal/constants"
	"net/http"
	"time"
)

type LoginMiddlewareBuilder struct {
	paths []string
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}

func (l *LoginMiddlewareBuilder) IgnorePaths(paths ...string) *LoginMiddlewareBuilder {
	for _, path := range paths {
		l.paths = append(l.paths, path)
	}
	return l
}

func (l *LoginMiddlewareBuilder) Build() gin.HandlerFunc {
	gob.Register(time.Now())
	return func(ctx *gin.Context) {
		for _, path := range l.paths {
			if ctx.Request.URL.Path == path {
				return
			}
		}
		session := sessions.Default(ctx)
		id := session.Get(constants.USER_ID)
		fmt.Println("id:", id)
		if id == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		now := time.Now()
		updateTime := session.Get("update_key")
		fmt.Println("updateTime:", updateTime)
		if updateTime == nil || now.Sub(updateTime.(time.Time)) > time.Second*2 {
			session.Set("update_key", now)
			//session.Set(constants.USER_ID, id)
			session.Options(sessions.Options{
				MaxAge: 10,
			})
			session.Save()
		}
	}
}
