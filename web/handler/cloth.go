package handler

import (
	"encoding/json"
	"fmt"
	"github.com/kons16/team7-backend/usecase"
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
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if method == "POST" {
		err := ch.cu.BuyCloth("sessionID", 123)
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
