package handler

import (
	"encoding/json"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	// メソッドのバリデート
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 返すjson
	ans := map[string]string{
		"message": "hello",
	}
	res, err := json.Marshal(ans)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
