package handler

import (
	"encoding/json"
	"fmt"
	"github.com/kons16/team7-backend/domain/entity"
	"github.com/kons16/team7-backend/usecase"
	"io/ioutil"
	"net/http"
	"strconv"
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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if method == "POST" {
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}

		cookie, err := r.Cookie("sessionID")

		var postData map[string]interface{}
		err = json.Unmarshal(body, &postData)
		if err != nil {
			fmt.Println(err)
		}

		var cordinate entity.Cordinate
		cordinate.Title = postData["title"].(string)
		cordinate.TopClothID, _ = strconv.Atoi(postData["top_cloth_id"].(string))
		cordinate.PantClothID, _ = strconv.Atoi(postData["pant_cloth_id"].(string))
		sessionID := cookie.Value

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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if method == "GET" {
		cookie, err := r.Cookie("sessionID")
		if err != nil {
			fmt.Println(err)
			return
		}
		v := cookie.Value

		cordinates := cdh.cdu.GetAll(v)
		// TopClothIDs, PantClothIDs から, その服情報を取得する
		var topIDs []int
		var pantIDs []int
		for _, v := range *cordinates {
			topIDs = append(topIDs, v.TopClothID)
			pantIDs = append(pantIDs, v.TopClothID)
		}

		topClothes := cdh.cdu.GetClothes(topIDs)
		pantClothes := cdh.cdu.GetClothes(topIDs)
		var tops []interface{}  // トップの詳細情報を格納
		var pants []interface{} // パンツの詳細情報を格納

		for _, v := range *topClothes {
			m := map[string]string{}
			m["ID"] = strconv.Itoa(v.ID)
			m["Name"] = v.Name
			m["Price"] = v.Price
			m["ImageUrl"] = v.ImageUrl
			m["Type"] = v.Type
			tops = append(tops, m)
		}

		for _, v := range *pantClothes {
			m := map[string]string{}
			m["ID"] = strconv.Itoa(v.ID)
			m["Name"] = v.Name
			m["Price"] = v.Price
			m["ImageUrl"] = v.ImageUrl
			m["Type"] = v.Type
			pants = append(pants, m)
		}

		// レスポンスデータ
		var s []interface{}
		for i, v := range *cordinates {
			m := map[string]string{}
			m["ID"] = strconv.Itoa(v.ID)
			m["Title"] = v.Title
			m["Top"] = tops[i].(string)
			m["Pant"] = pants[i].(string)
			s = append(s, m)
		}

		ans := map[string]interface{}{
			"clothes": s,
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
