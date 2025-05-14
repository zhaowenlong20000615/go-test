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
