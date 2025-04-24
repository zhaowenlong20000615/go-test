package controller

import (
	"fmt"
	"go-test/go-blog/dao"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(dao.DB)
	w.Write([]byte("Hello World!"))
}
