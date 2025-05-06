package controller

import (
	"go-test/go-blog/common"
	"go-test/go-blog/config"
	"net/http"
)

func RegisterHtml(w http.ResponseWriter, r *http.Request) {
	registerTemplate := common.Template.Register
	registerTemplate.WriteData(w, config.Config.Viewer)
}
