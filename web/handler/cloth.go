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
	w.Header().Set("Access-Control-Allow-Origin", "http")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

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

	if method == "GET" {
		clothes := ch.cu.GetAll()
		clothesJson, err := json.Marshal(clothes)

		ans := map[string]string{
			"clothes": string(clothesJson),
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
