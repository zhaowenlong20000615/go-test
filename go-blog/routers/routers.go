package routers

import (
	"go-test/go-blog/controller"
	"net/http"
)

func Routes() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/login", controller.LoginHtml)
	http.HandleFunc("/register", controller.RegisterHtml)
	http.HandleFunc("/writing", controller.WriteHtml)
	http.HandleFunc("/api/v1/login", controller.Login)
	http.HandleFunc("/api/v1/register", controller.Register)
	http.HandleFunc("/api/v1/post", controller.AddOrUpdate)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource"))))
}
