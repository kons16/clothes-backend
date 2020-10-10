package handler

import (
	"encoding/json"
	"fmt"
	"github.com/kons16/team7-backend/usecase"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type SessionHandler struct {
	sc *usecase.SessionUseCase
}

func NewSessionHandler(sessionUseCase *usecase.SessionUseCase) *SessionHandler {
	return &SessionHandler{sc: sessionUseCase}
}

// POST /login submit_id と password からログインし、sessionID を返す
func (sh *SessionHandler) Login(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	fmt.Println("[method] " + method)
	for k, v := range r.Header {
		fmt.Print("[header] " + k)
		fmt.Println(": " + strings.Join(v, ","))
	}

	if method == "POST" {
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		var postData map[string]interface{}
		err = json.Unmarshal(body, &postData)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(postData)

		var userLogin usecase.UserLogin
		userLogin.SubmitID = postData["submit_id"].(string)
		userLogin.Password = postData["password"].(string)

		_, err = sh.sc.Login(&userLogin)
	}
}

// GET /is_login cookieから受け取った sessionID を元にユーザーがログインしているかどうか確認
func (sh *SessionHandler) FindUserBySession(w http.ResponseWriter, r *http.Request) {

}
