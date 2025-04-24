package common

import (
	"go-test/go-blog/config"
	"go-test/go-blog/models"
	"sync"
)

var Template models.HtmlTemplate

func Load() {
	var err error

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		path := config.Config.System.CurrentDir + "\\template"
		Template, err = models.InitHtmlTemplate(path)
		if err != nil {
			panic(err)
		}
		wg.Done()
	}()
	wg.Wait()
}
