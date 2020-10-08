package usecase

import "github.com/kons16/team7-backend/domain/repository"

type SessionUseCase struct {
	sessionRepo repository.Session
}

func NewSessionUseCase(sessionRepo repository.Session) *SessionUseCase {
	return &SessionUseCase{sessionRepo: sessionRepo}
}
