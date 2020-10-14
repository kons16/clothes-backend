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
	userID := sc.sessionRepo.CheckBySession(sessionID)
	return userID
}

func (sc *SessionUseCase) Logout(sessionID string) bool {
	check := sc.sessionRepo.Logout(sessionID)
	return check
}
