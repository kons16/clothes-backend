package handler

import (
	"encoding/json"
	"fmt"
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
	method := r.Method
	fmt.Println("[method] " + method)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if method == "GET" {
		cookie, err := r.Cookie("sessionID")
		if err != nil {
			fmt.Println(err)
			return
		}
		v := cookie.Value

		getUserID := sh.sc.CheckBySession(v)
		checkStr := ""
		// getUserID があれば true に変更
		if getUserID != 0 {
			checkStr = "true"
		} else {
			checkStr = "false"
		}

		// isLogin が true のとき session が残っている
		ans := map[string]string{
			"is_login": checkStr,
		}
		res, err := json.Marshal(ans)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
}

// GET /logout Redis から SessionID の カラムを削除する. 削除できたら true を受け取る
func (sh *SessionHandler) Logout(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	fmt.Println("[method] " + method)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if method == "GET" {
		cookie, err := r.Cookie("sessionID")
		if err != nil {
			fmt.Println(err)
			return
		}
		v := cookie.Value

		checkBool := sh.sc.Logout(v)
		checkStr := ""
		if err != nil {
			fmt.Println(err)
		}
		// bool を str に変換
		if checkBool == true {
			checkStr = "true"
		} else {
			checkStr = "false"
		}

		// logout が true のとき logout 成功
		ans := map[string]string{
			"is_logout": checkStr,
		}
		res, err := json.Marshal(ans)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
}
