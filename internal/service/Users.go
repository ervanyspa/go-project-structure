package service

import (
	"context"
	"project-structure/internal/model"
	"project-structure/internal/repository"
)

type UserService interface {
	GetUsers(ctx context.Context) ([]model.User, error)
	GetUserByID(ctx context.Context, id uint64) (model.User, error)
}

type userServiceImpl struct{
	repo repository.UserQuery
}

func NewUserService(repo repository.UserQuery) UserService{
	return &userServiceImpl{repo: repo}

}

func (u *userServiceImpl) GetUsers(ctx context.Context) ([]model.User, error) {
	users, err := u.repo.GetUsers(ctx)
	if err != nil {
		return nil,err
	}
	return users, err
}

func (u *userServiceImpl) GetUserByID(ctx context.Context, id uint64) (model.User, error) {
	user, err := u.repo.GetUserByID(ctx, id)
	if err != nil {
		return model.User{}, err
	}
	return user, err
}


