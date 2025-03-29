package web

import "github.com/gin-gonic/gin"

func RegisterRouters() *gin.Engine {
	server := gin.Default()
	return server
}
