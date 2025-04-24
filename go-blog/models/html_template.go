package models

import (
	"html/template"
	"time"
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
	htmlTemplate.Index = tp[0]
	htmlTemplate.Category = tp[1]
	htmlTemplate.Pigeonhole = tp[2]
	htmlTemplate.Login = tp[3]
	htmlTemplate.Detail = tp[4]
	htmlTemplate.Writing = tp[5]
	return htmlTemplate, nil
}

func readHtmlTemplate(htmlFilename []string, viewDir string) ([]TemplatePointer, error) {
	var htmlTemplate []TemplatePointer

	head := viewDir + "/layout/header.html"
	footer := viewDir + "/layout/footer.html"
	home := viewDir + "/home.html"
	personal := viewDir + "/layout/personal.html"
	postList := viewDir + "/layout/post-list.html"
	pagination := viewDir + "/layout/pagination.html"

	for _, name := range htmlFilename {
		filePath := viewDir + "/" + name + ".html"
		tp := template.New(name + ".html")
		tp.Funcs(template.FuncMap{"isODD": IsODD, "date": Date, "dateDay": DateDay, "getNextName": GetNextName})
		var err error
		tp, err = tp.ParseFiles(filePath, head, footer, home, personal, postList, pagination)
		if err != nil {
			return htmlTemplate, err
		}
		htmlTemplate = append(htmlTemplate, TemplatePointer{tp})
	}
	return htmlTemplate, nil
}

//	func SpreadDigit(n int) []int {
//		var r []int
//		for i := 1; i <= n; i++ {
//			r = append(r, i)
//		}
//		return r
//	}
func IsODD(num int) bool {
	return num%2 == 0
}
func Date(layout string) string {
	return time.Now().Format(layout)
}
func DateDay(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}
func GetNextName(strs []string, i int) interface{} {
	return strs[i+1]
}
