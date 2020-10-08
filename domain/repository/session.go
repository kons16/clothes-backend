package repository

import "github.com/kons16/team7-backend/domain/entity"

type Session interface {
	FindUserBySession(sessionID int64) (*entity.User, error)
	CreateUserSession() (int64, error)
}
