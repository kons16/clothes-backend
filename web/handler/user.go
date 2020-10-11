package handler

import (
	"encoding/json"
	"fmt"
	"github.com/kons16/team7-backend/usecase"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type UserHandler struct {
	uc *usecase.UserUseCase
}

func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{uc: userUseCase}
}

// POST /user ユーザーを新規登録する
func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	fmt.Println("[method] " + method)

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if method == "POST" {
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}

		var postData map[string]interface{}
		err = json.Unmarshal(body, &postData)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(postData)

		var user usecase.User
		user.Name = postData["name"].(string)
		user.SubmitID = postData["submit_id"].(string)
		user.Password = postData["password"].(string)
		user.Year, _ = strconv.Atoi(postData["year"].(string))
		user.Sex, _ = strconv.Atoi(postData["sex"].(string))

		sessionID, err := uh.uc.CreateUser(&user)
		if err != nil {
			fmt.Println(err)
		}

		ans := map[string]string{
			"sessionID": sessionID,
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

// POST /login submit_id と password からログインし、sessionID を返す
func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	fmt.Println("[method] " + method)

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
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
			fmt.Println(err)
		}

		fmt.Println(postData)

		var userLogin usecase.UserLogin
		userLogin.SubmitID = postData["submit_id"].(string)
		userLogin.Password = postData["password"].(string)

		sessionID, err := uh.uc.Login(&userLogin)
		if err != nil {
			fmt.Println(err)
		}

		ans := map[string]string{
			"sessionID": sessionID,
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
