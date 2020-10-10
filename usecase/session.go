package usecase

import (
	"github.com/kons16/team7-backend/domain/repository"
)

type SessionUseCase struct {
	sessionRepo repository.Session
}

// ログインする際に使用する構造体
type UserLogin struct {
	SubmitID string
	Password string
}

func NewSessionUseCase(sessionRepo repository.Session) *SessionUseCase {
	return &SessionUseCase{sessionRepo: sessionRepo}
}

func (su *SessionUseCase) Login(userLogin *UserLogin) (string, error) {
	// submit_id に紐づく password_hash を取得する

}
