package handler

import (
	"encoding/json"
	"fmt"
	"github.com/kons16/team7-backend/usecase"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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

		var user usecase.User
		user.Name = postData["name"].(string)
		user.SubmitID = postData["submit_id"].(string)
		user.Password = postData["password"].(string)
		user.Year, _ = strconv.Atoi(postData["year"].(string))
		user.Sex, _ = strconv.Atoi(postData["sex"].(string))

		uh.uc.CreateUser(&user)
	}
}
