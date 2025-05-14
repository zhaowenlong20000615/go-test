package common

import (
	"encoding/json"
	"go-test/go-blog/config"
	"go-test/go-blog/models"
	"net/http"
	"sync"
)

var Template models.HtmlTemplate

func Load() {
	var err error

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		path := config.Config.System.CurrentDir + "/template"
		Template, err = models.InitHtmlTemplate(path)
		if err != nil {
			panic(err)
		}
		wg.Done()
	}()
	wg.Wait()
}

func Error(w http.ResponseWriter, err error) {
	var ret = models.Result{}
	ret.Code = -999
	ret.Msg = err.Error()
	ret.Data = nil
	r, err := json.Marshal(ret)
	if err != nil {
		ret.Msg = err.Error()
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(r)
}

func Success(w http.ResponseWriter, data interface{}) {
	var ret = models.Result{}
	ret.Code = 200
	ret.Msg = "success"
	ret.Data = data
	r, err := json.Marshal(ret)
	if err != nil {
		ret.Msg = err.Error()
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(r)
}
