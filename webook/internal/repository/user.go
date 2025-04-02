package repository

import (
	"context"
	"go-test/webook/internal/domain"
	"go-test/webook/internal/repository/dao"
	"time"
)

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
		NickName: u.NickName,
		Profile:  u.Profile,
		BirthDay: u.BirthDay,
	})
}

func (r *UserRepository) FindByEmail(ctx context.Context, u domain.User) (domain.User, error) {
	user, err := r.dao.FindByEmail(ctx, u.Email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Id:       user.Id,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id int64) (domain.User, error) {
	user, err := r.dao.FindByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Id:       user.Id,
		Email:    user.Email,
		Password: user.Password,
		Ctime:    time.UnixMilli(user.Ctime).Format("2006-01-02 15:04:05"),
		Utime:    time.UnixMilli(user.Utime).Format("2006-01-02 15:04:05"),
	}, nil
}

func (r *UserRepository) DeleteByID(ctx context.Context, id int64) error {
	return r.dao.DeleteByID(ctx, id)
}

func (r *UserRepository) UpdateByID(ctx context.Context, id int64, u domain.User) error {
	return r.dao.UpdateByID(ctx, id, dao.User{
		NickName: u.NickName,
		Profile:  u.Profile,
		BirthDay: u.BirthDay,
	})
}
