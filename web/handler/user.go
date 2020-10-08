package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// POST /user ユーザーを新規登録する
func CreateUser(w http.ResponseWriter, r *http.Request) {
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
		fmt.Println(postData["Name"])
	}
}
