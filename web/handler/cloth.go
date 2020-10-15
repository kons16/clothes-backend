package handler

import (
	"encoding/json"
	"fmt"
	"github.com/kons16/team7-backend/usecase"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ClothHandler struct {
	cu *usecase.ClothUseCase
}

func NewClothHandler(clothUseCase *usecase.ClothUseCase) *ClothHandler {
	return &ClothHandler{cu: clothUseCase}
}

// POST /post_cloth 新しくclothesテーブルに服を追加する
func (ch *ClothHandler) CreateCloth(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	fmt.Println("[method] " + method)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if method == "POST" {
		defer r.Body.Close()

		var cloth usecase.Cloth
		cloth.Name = r.FormValue("name")
		cloth.Price = r.FormValue("price")
		cloth.Type = r.FormValue("type")
		if r.FormValue("image") == "data:," {
			cloth.ImageBase64 = ""
		} else {
			cloth.ImageBase64 = r.FormValue("image")
		}

		clothID, err := ch.cu.CreateCloth(&cloth)
		if err != nil {
			fmt.Println(err)
			return
		}

		ans := map[string]string{
			"cloth_id": strconv.Itoa(clothID),
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

// GET /get_cloth 服情報をすべて取得する
func (ch *ClothHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	fmt.Println("[method] " + method)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if method == "GET" {
		clothes := ch.cu.GetAll()
		var s []interface{}
		for _, v := range *clothes {
			m := map[string]string{}
			m["ID"] = strconv.Itoa(v.ID)
			m["Name"] = v.Name
			m["Price"] = v.Price
			m["ImageUrl"] = v.ImageUrl
			m["Type"] = v.Type
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

// POST /buy 受け取った服の ID より服を購入する
func (ch *ClothHandler) BuyCloth(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	fmt.Println("[method] " + method)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if method == "POST" {
		cookie, err := r.Cookie("sessionID")
		if err != nil {
			fmt.Println(err)
			return
		}
		v := cookie.Value

		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		var postData map[string]interface{}
		err = json.Unmarshal(body, &postData)
		if err != nil {
			fmt.Println(err)
			return
		}

		clothID, _ := strconv.Atoi(postData["clothID"].(string))

		err = ch.cu.BuyCloth(v, clothID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ans := map[string]interface{}{
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

// GET /my_cloth 購入した服の情報を取得する
func (ch *ClothHandler) GetBuyCloth(w http.ResponseWriter, r *http.Request) {
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
		clothes := ch.cu.GetBuyCloth(v)

		var s []interface{}
		for _, v := range *clothes {
			m := map[string]string{}
			m["ID"] = strconv.Itoa(v.ID)
			m["Name"] = v.Name
			m["Price"] = v.Price
			m["ImageUrl"] = v.ImageUrl
			m["Type"] = v.Type
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
