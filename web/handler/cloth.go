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

		var cloth usecase.Cloth
		cloth.Name = postData["name"].(string)
		cloth.Price = postData["price"].(string)
		cloth.ImageBase64 = postData["image"].(string)

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
