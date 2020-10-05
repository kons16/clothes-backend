package web

import (
	"github.com/kons16/team7-backend/usecase"
	"net/http"
)

func NewServer(userUC *usecase.UserUseCase) *http.Server {
	s := http.Server{
		Addr: "127.0.0.1:8000",
	}
	return &s
}
