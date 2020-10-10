package web

import (
	"github.com/kons16/team7-backend/usecase"
	"github.com/kons16/team7-backend/web/handler"
	"net/http"
)

func NewServer(userUC *usecase.UserUseCase, sessionUC *usecase.SessionUseCase) *http.Server {
	var s http.Server
	s.Addr = ":8000"

	userHandler := handler.NewUserHandler(userUC)
	sessionHandler := handler.NewSessionHandler(sessionUC)

	// GET サーバーが立ち上がっているか確認
	http.HandleFunc("/hello", handler.Hello)
	// POST アカウントの新規登録し, sessionID を返す
	http.HandleFunc("/user", userHandler.CreateUser)
	// GET cookie の sessionID から該当するユーザー情報を返す
	http.HandleFunc("/is_login", sessionHandler.FindUserBySession)
	// POST user_sessionテーブル に userID に紐づく sessionID のカラムを追加し,sessionID を返す
	http.HandleFunc("/login", sessionHandler.Login)
	// GET user_sessionテーブル から sessionID のカラムを削除する
	http.HandleFunc("/logout", sessionHandler.FindUserBySession)
	// POST 購入した服の ID を ユーザーと紐付ける
	// http.HandleFunc("/buy_cloth", nil)
	// GET 購入した服の情報を持ってくる
	// http.HandleFunc("/cloth", nil)

	return &s
}
