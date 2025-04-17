package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", index)

	println("启动服务在：127.0.0.1:8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("http server listen err: %v", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	var indexData IndexData
	indexData.Title = "标题"
	indexData.Desc = "描述"
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	//访问博客首页模板的时候，因为有多个模板的嵌套，解析文件的时候，需要将其涉及到的所有模板都进行解析
	home := dir + "/template/home.html"
	header := dir + "/template/layout/header.html"
	footer := dir + "/template/layout/footer.html"
	personal := dir + "/template/layout/personal.html"
	post := dir + "/template/layout/post-list.html"
	pagination := dir + "/template/layout/pagination.html"
	files, err := template.New("index.html").ParseFiles(dir+"/template/index.html", home, header, footer, personal, post, pagination)
	if err != nil {
		fmt.Println(err)
	}
	err = files.Execute(w, indexData)
	fmt.Println(err)
}
