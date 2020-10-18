package handler

import (
	"fmt"
	"github.com/kons16/team7-backend/usecase"
	"net/http"
)

type CordinateHandler struct {
	cdu *usecase.CordinateUseCase
}

func NewCordinateHandler(cordinateUseCase *usecase.CordinateUseCase) *CordinateHandler {
	return &CordinateHandler{cdu: cordinateUseCase}
}

// POST /cordinate 新しく cordinateテーブル にコーディネートを追加する
func (cdh *CordinateHandler) CreateCordinate(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	fmt.Println("[method] " + method)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if method == "POST" {
	}
}

// GET /get_cordinate ログインしているユーザーのコーディネートを全て取得する
func (cdh *CordinateHandler) Get(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	fmt.Println("[method] " + method)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if method == "GET" {
	}
}
