package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	serve := gin.Default()

	serve.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world")
	})

	serve.GET("/user/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.String(http.StatusOK, "params 参数 User"+id)
	})

	serve.GET("/user", func(ctx *gin.Context) {
		query := ctx.Request.URL.Query()
		for k, v := range query {
			fmt.Println(k, v)
		}
		println(query.Encode())
		ctx.String(http.StatusOK, "查询参数 User: "+query.Encode())
	})

	serve.GET("/test/*.html", func(ctx *gin.Context) {
		params := ctx.Params
		println(params.ByName(".html"))
		ctx.String(http.StatusOK, "通配符路由："+params.ByName(".html"))
	})

	//go server.Run(":8081")

	serve.Run(":8080")
	println("Hello Gin")
}
