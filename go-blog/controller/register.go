package controller

import (
	"encoding/json"
	"fmt"
	"go-test/go-blog/common"
	"go-test/go-blog/config"
	"go-test/go-blog/dao"
	"go-test/go-blog/models"
	"go-test/go-blog/utils"
	"net/http"
)

func RegisterHtml(w http.ResponseWriter, r *http.Request) {
	registerTemplate := common.Template.Register
	registerTemplate.WriteData(w, config.Config.Viewer)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var params = models.RegisterReq{}
	json.NewDecoder(r.Body).Decode(&params)
	fmt.Printf("params: %#v\n", params)
	params.Passwd = utils.Md5Crypt(params.Passwd)
	user, err := dao.Register(params)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, user)
}
