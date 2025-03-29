package repository

import (
	"context"
	"go-test/webook/internal/domain"
	"go-test/webook/internal/repository/dao"
)

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (r *UserRepository) create(ctx context.Context, u domain.User) error {
	return nil
}
