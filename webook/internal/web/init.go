package web

import "github.com/gin-gonic/gin"

func RegisterRouters(serve *gin.Engine) {
	RegisterUser(serve)
}
