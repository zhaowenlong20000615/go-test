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
	"time"
)

func LoginHtml(w http.ResponseWriter, r *http.Request) {
	loginTemplate := common.Template.Login
	loginTemplate.WriteData(w, config.Config.Viewer)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var params = models.LoginReq{}
	all, err := ioutil.ReadAll(r.Body)
	if err != nil {
		common.Error(w, err)
		return
	}
	json.Unmarshal(all, &params)
	params.Passwd = utils.Md5Crypt(params.Passwd)
	user, err := dao.Login(params)
	if err != nil {
		common.Error(w, err)
		return
	}
	token, err := utils.CrateToken(user, time.Hour*24*7)
	if err != nil {
		common.Error(w, err)
		return
	}
	var ret = models.LoginRes{UserInfo: user, Token: token}
	common.Success(w, ret)
}
