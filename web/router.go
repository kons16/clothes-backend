package web

import (
	"github.com/kons16/team7-backend/usecase"
	"github.com/kons16/team7-backend/web/handler"
	"net/http"
)

func NewServer(userUC *usecase.UserUseCase, sessionUC *usecase.SessionUseCase) *http.Server {
	var s http.Server
	s.Addr = ":8000"

	http.HandleFunc("/hello", handler.Hello)
	http.HandleFunc("/user", handler.CreateUser)
	http.HandleFunc("/is_login", handler.FindUserBySession)

	return &s
}
