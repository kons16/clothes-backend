package handler

import (
	"github.com/kons16/team7-backend/usecase"
	"net/http"
)

type SessionHandler struct {
	sc *usecase.SessionUseCase
}

func NewSessionHandler(sessionUseCase *usecase.SessionUseCase) *SessionHandler {
	return &SessionHandler{sc: sessionUseCase}
}

// GET /is_login cookieから受け取った sessionID を元にユーザーがログインしているかどうか確認
func (sh *SessionHandler) FindUserBySession(w http.ResponseWriter, r *http.Request) {

}
