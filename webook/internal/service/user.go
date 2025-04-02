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
	return user, nil
}

func (svg *UserService) GetUserInfo(ctx context.Context, id int64) (domain.User, error) {
	user, err := svg.repo.FindByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (svg *UserService) DeleteUser(ctx context.Context, id int64) error {
	_, err := svg.repo.FindByID(ctx, id)
	if err != nil {
		return constants.ErrNotFoundUser
	}
	return svg.repo.DeleteByID(ctx, id)
}

func (svg *UserService) UpdateUser(ctx context.Context, id int64, u domain.User) error {
	err := svg.repo.UpdateByID(ctx, id, u)
	return err
}
