package controller

import (
	"go-test/go-blog/common"
	"go-test/go-blog/config"
	"go-test/go-blog/dao"
	"net/http"
)

func WriteHtml(w http.ResponseWriter, r *http.Request) {
	writingHtml := common.Template.Writing
	m := make(map[string]interface{})
	m["categorys"] = dao.GetCategorys()
	m["CdnURL"] = config.Config.System.CdnURL
	m["Title"] = config.Config.Viewer.Title
	writingHtml.WriteData(w, m)
}
