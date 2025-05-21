package models

import (
	"html/template"
)

type Post struct {
	Pid        int    `json:"pid"`        // 文章ID
	Title      string `json:"title"`      // 文章ID
	Slug       string `json:"slug"`       // 自定也页面 path
	Content    string `json:"content"`    // 文章的html
	Markdown   string `json:"markdown"`   // 文章的Markdown
	CategoryId int    `json:"categoryId"` //分类id
	UserId     int    `json:"userId"`     //用户id
	ViewCount  int    `json:"viewCount"`  //查看次数
	Type       int    `json:"type"`       //文章类型 0 普通，1 自定义文章
	CreateAt   int64  `json:"createAt"`   // 创建时间
	UpdateAt   int64  `json:"updateAt"`   // 更新时间
}

type PostReq struct {
	Title      string `json:"title"`      // 文章ID
	Slug       string `json:"slug"`       // 自定也页面 path
	Content    string `json:"content"`    // 文章的html
	Markdown   string `json:"markdown"`   // 文章的Markdown
	CategoryId int    `json:"categoryId"` //分类id
	UserId     int    `json:"userId"`     //用户id
	ViewCount  int    `json:"viewCount"`  //查看次数
	Type       int    `json:"type"`       //文章类型 0 普通，1 自定义文章
}

type PostMore struct {
	Pid          int           `json:"pid"`          // 文章ID
	Title        string        `json:"title"`        // 文章ID
	Slug         string        `json:"slug"`         // 自定也页面 path
	Content      template.HTML `json:"content"`      // 文章的html
	CategoryId   int           `json:"categoryId"`   // 文章的Markdown
	CategoryName string        `json:"categoryName"` // 分类名
	UserId       int           `json:"userId"`       // 用户id
	UserName     string        `json:"userName"`     // 用户名
	ViewCount    int           `json:"viewCount"`    // 查看次数
	Type         int           `json:"type"`         // 文章类型 0 普通，1 自定义文章
	CreateAt     string        `json:"createAt"`
	UpdateAt     string        `json:"updateAt"`
}

type PostQuery struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}
