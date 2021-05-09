package repository

import "github.com/kons16/team7-backend/domain/entity"

type User interface {
	FindByID(id string) (*entity.User, error)
	Create(user *entity.User) (int, error)
	FindUserBySubmitID(submitID string) (*entity.LoginGetUser, error)
}
