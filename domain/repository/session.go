package repository

import "github.com/kons16/team7-backend/domain/entity"

type Session interface {
	FindUserBySession(sessionID int) (*entity.User, error)
	CreateUserSession(id int) (string, error)
}
