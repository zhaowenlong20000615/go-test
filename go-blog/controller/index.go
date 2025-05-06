package controller

import (
	"go-test/go-blog/common"
	"go-test/go-blog/config"
	"go-test/go-blog/dao"
	"go-test/go-blog/models"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	indexTemplate := common.Template.Index
	categorys := dao.GetCategorys()

	homeData := models.HomeData{
		Viewer:    config.Config.Viewer,
		Categorys: categorys,
		Posts:     nil,
		Total:     0,
		Page:      0,
		Pages:     nil,
		PageEnd:   false,
	}
	indexTemplate.WriteData(w, homeData)
}
