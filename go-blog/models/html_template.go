package models

import (
	"fmt"
	"html/template"
)

type TemplatePointer struct {
	*template.Template
}

type HtmlTemplate struct {
	Index      TemplatePointer
	Category   TemplatePointer
	Pigeonhole TemplatePointer
	Login      TemplatePointer
	Detail     TemplatePointer
	Writing    TemplatePointer
}

func InitHtmlTemplate(viewDir string) (HtmlTemplate, error) {
	var htmlTemplate HtmlTemplate
	tp, err := readHtmlTemplate([]string{"index", "category", "pigeonhole", "login", "detail", "writing"}, viewDir)
	if err != nil {
		return htmlTemplate, err
	}
	fmt.Println("InitHtmlTemplate", tp)
	return htmlTemplate, nil
}

func readHtmlTemplate(htmlFilename []string, viewDir string) ([]TemplatePointer, error) {
	var htmlTemplate []TemplatePointer
	for _, name := range htmlFilename {
		filePath := viewDir + "/" + name + ".html"
		fmt.Println(filePath)
	}
	return htmlTemplate, nil
}
