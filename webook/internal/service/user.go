package service

import (
	"context"
	"go-test/webook/internal/domain"
	"go-test/webook/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) AddUser(ctx context.Context, u domain.User) error {
	return svc.repo.Create(ctx, u)
}
