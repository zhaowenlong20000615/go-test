package routers

import (
	"go-test/go-blog/controller"
	"net/http"
)

func Routes() {
	http.HandleFunc("/", controller.Index)
}
