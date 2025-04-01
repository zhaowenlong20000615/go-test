package service

import (
	"context"
	"go-test/webook/internal/constants"
	"go-test/webook/internal/domain"
	"go-test/webook/internal/repository"
	"golang.org/x/crypto/bcrypt"
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
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return svc.repo.Create(ctx, u)
}

func (svc *UserService) Login(ctx context.Context, u domain.User) (domain.User, error) {
	user, err := svc.repo.FindByEmail(ctx, u)
	if err != nil {
		return domain.User{}, constants.ErrNotFoundUser
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		return domain.User{}, constants.ErrInvaildUserOrPassword
	}
	return u, nil
}
