package dao

import (
	"go-test/go-blog/models"
)

func Register(req models.RegisterReq) (models.User, error) {
	ret, err := DB.Exec("INSERT INTO `user` (username, passwd) VALUES (?, ?)", req.Name, req.Passwd)
	if err != nil {
		return models.User{}, err
	}

	id, err := ret.LastInsertId()
	if err != nil {
		return models.User{}, err
	}
	return models.User{Uid: int(id)}, nil
}

func Login(req models.LoginReq) (models.User, error) {
	row := DB.QueryRow("SELECT * FROM user WHERE username=? and passwd=? LIMIT 1", req.Name, req.Passwd)
	var user models.User
	err := row.Scan(&user.Uid, &user.UserName, &user.Passwd, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
