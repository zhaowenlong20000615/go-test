package controller

import (
	"encoding/json"
	"go-test/go-blog/common"
	"go-test/go-blog/config"
	"go-test/go-blog/dao"
	"go-test/go-blog/models"
	"go-test/go-blog/utils"
	"io/ioutil"
	"net/http"
	"strconv"
)

func WriteHtml(w http.ResponseWriter, r *http.Request) {
	writingHtml := common.Template.Writing
	m := make(map[string]interface{})
	m["categorys"] = dao.GetCategorys()
	m["CdnURL"] = config.Config.System.CdnURL
	m["Title"] = config.Config.Viewer.Title
	writingHtml.WriteData(w, m)
}

func AddOrUpdate(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "POST" {
		addArticle(w, r)
		return
	}
	if method == "PUT" {
		updateArticle(w, r)
		return
	}
}

func addArticle(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("AUTHORIZATION")
	claims, err := utils.ParseToken(token)
	if err != nil {
		common.Error(w, err)
		return
	}
	post := models.PostReq{}
	post.UserId = claims.User.Uid
	all, err := ioutil.ReadAll(r.Body)
	if err != nil {
		common.Error(w, err)
	}
	ts := struct {
		CategoryId string `json:"categoryId"`
	}{}
	json.Unmarshal(all, &ts)
	json.Unmarshal(all, &post)
	atoi, err := strconv.Atoi(ts.CategoryId)
	post.CategoryId = atoi
	savePost, err := dao.SavePost(post)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, savePost)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {

}
