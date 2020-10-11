package handler

import (
	"encoding/json"
	"fmt"
	"github.com/kons16/team7-backend/usecase"
	"net/http"
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
		cookie, err := r.Cookie("sessionID")
		if err != nil {
			fmt.Println(err)
			return
		}
		v := cookie.Value

		checkBool := sh.sc.CheckBySession(v)
		checkStr := ""
		if err != nil {
			fmt.Println(err)
		}
		// bool を str に変換
		if checkBool == true {
			checkStr = "true"
		} else {
			checkStr = "false"
		}

		// isLogin が true のとき session が残っている
		ans := map[string]string{
			"cloth_id": checkStr,
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
