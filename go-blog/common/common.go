package common

import (
	"fmt"
	"go-test/go-blog/models"
)

func Load() {
	fmt.Println()
	template, err := models.InitHtmlTemplate("")
	if err != nil {
		panic(err)
	}
	fmt.Println(template)
}
