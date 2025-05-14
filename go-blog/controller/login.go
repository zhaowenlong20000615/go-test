package controller

import (
	"encoding/json"
	"fmt"
	"go-test/go-blog/common"
	"go-test/go-blog/config"
	"go-test/go-blog/models"
	"io/ioutil"
	"net/http"
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
	fmt.Printf("params: %#v\n", params)
}
