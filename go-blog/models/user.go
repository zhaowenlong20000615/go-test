package models

import "time"

type User struct {
	Uid      int       `json:"uid"`      // 用户id
	UserName string    `json:"userName"` // 用户名
	Avatar   string    `json:"avatar"`   // 头像
	Passwd   string    `json:"-"`        // 密码
	CreateAt time.Time `json:"createAt"` // 创建时间
	UpdateAt time.Time `json:"updateAt"` // 创建时间
}

type RegisterReq struct {
	Name   string `json:"username"`
	Passwd string `p:"passwd" v:"required|length:6,255#请输入密码|密码长度不够"`
}
type LoginReq struct {
	Name   string `json:"username"`
	Passwd string `json:"passwd"`
}
