package main

import (
	"go-test/go-blog/common"
	"go-test/go-blog/server"
)

func init() {
	common.Load()
}

func main() {
	server.App.Start("127.0.0.1", "8080")
}
