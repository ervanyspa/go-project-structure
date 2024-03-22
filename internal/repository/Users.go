package repository

import (
	"context"
	"project-structure/internal/infrastructure"
	"project-structure/internal/model"
)

type UserQuery interface {
	GetUsers(ctx context.Context) ([]model.User, error)
	GetUserByID(ctx context.Context, id uint64) (model.User, error)
}

type UserCommand interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
}

type userQueryImpl struct{
	db infrastructure.GromPostgres
}

func NewUserQuery(db infrastructure.GromPostgres) UserQuery {
	return &userQueryImpl{db: db}
}

func (u *userQueryImpl) GetUsers(ctx context.Context) ([]model.User, error) {
	db := u.db.GetConnection()
	users := []model.User{}
	if err := db.WithContext(ctx).Table("users").Find(&users).Error; err != nil {
		return nil, nil
	}
	return users, nil
}

func (u *userQueryImpl) GetUserByID(ctx context.Context, id uint64) (model.User, error) {
	db := u.db.GetConnection()
	users := model.User{}
	if err := db.WithContext(ctx).Table("users").Where("id = ?", id).Find(&users).Error; err != nil {
		return model.User{}, nil
	}
	return users, nil
}