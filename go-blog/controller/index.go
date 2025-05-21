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
	query := models.PostQuery{
		Page:     1,
		PageSize: 10,
	}
	posts, err := dao.GetPost(query)
	if err != nil {
		common.Error(w, err)
		return
	}
	total := len(posts)
	pagesLen := (total / query.PageSize) + 1
	var pages []int
	for i := 0; i < pagesLen; i++ {
		pages = append(pages, i+1)
	}
	homeData := models.HomeData{
		Viewer:    config.Config.Viewer,
		Categorys: categorys,
		Posts:     posts,
		Total:     total,
		Page:      query.Page,
		Pages:     pages,
		PageEnd:   query.Page != pagesLen,
	}
	indexTemplate.WriteData(w, homeData)
}
