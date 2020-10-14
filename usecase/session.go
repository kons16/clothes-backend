package usecase

import (
	"github.com/kons16/team7-backend/domain/repository"
)

type SessionUseCase struct {
	sessionRepo repository.Session
}

func NewSessionUseCase(sessionRepo repository.Session) *SessionUseCase {
	return &SessionUseCase{sessionRepo: sessionRepo}
}

func (sc *SessionUseCase) CheckBySession(sessionID string) int {
	check := sc.sessionRepo.CheckBySession(sessionID)
	return check
}

func (sc *SessionUseCase) Logout(sessionID string) bool {
	check := sc.sessionRepo.Logout(sessionID)
	return check
}
