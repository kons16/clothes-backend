package handler

import (
	"encoding/json"
	"fmt"
	"github.com/kons16/team7-backend/domain/entity"
	"github.com/kons16/team7-backend/usecase"
	"io/ioutil"
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

		var cordinate entity.Cordinate
		cordinate.Title = postData["title"].(string)
		cordinate.TopClothID = postData["top_cloth_id"].(string)
		cordinate.PantClothID = postData["pant_cloth_id"].(string)
		sessionID := postData["session_id"].(string)

		err = cdh.cdu.CreateCordinate(&cordinate, sessionID)
		if err != nil {
			fmt.Println(err)
		}

		ans := map[string]string{
			"message": "success",
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

// GET /get_cordinate ログインしているユーザーのコーディネートを全て取得する
func (cdh *CordinateHandler) Get(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	fmt.Println("[method] " + method)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if method == "GET" {
	}
}
