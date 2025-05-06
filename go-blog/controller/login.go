package controller

import (
	"go-test/go-blog/common"
	"go-test/go-blog/config"
	"net/http"
)

func LoginHtml(w http.ResponseWriter, r *http.Request) {
	loginTemplate := common.Template.Login
	loginTemplate.WriteData(w, config.Config.Viewer)
}
