package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-test/webook/internal/web"
	"strings"
	"time"
)

func main() {
	serve := web.RegisterRouters()
	serve.Use(func(context *gin.Context) {
		println("第一个中间件！")
	})
	serve.Use(func(context *gin.Context) {
		println("第二个中间件！")
	})
	serve.Use(cors.New(cors.Config{
		//AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "https://github.com")
		},
		MaxAge: 12 * time.Hour,
	}))
	serve.Run(":8080")
}
