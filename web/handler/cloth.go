package handler

import (
	"encoding/json"
	"fmt"
	"github.com/kons16/team7-backend/usecase"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
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
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}

		postData := map[string]string{}
		sBody := strings.Split(string(body), "&")
		for _, v := range sBody {
			postData[strings.Split(v, "=")[0]] = strings.Split(v, "=")[1]
		}

		var cloth usecase.Cloth
		cloth.Name = postData["name"]
		cloth.Price = postData["price"]
		cloth.Type = postData["type"]
		cloth.ImageBase64 = postData["image"]

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
