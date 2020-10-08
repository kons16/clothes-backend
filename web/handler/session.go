package handler

import "net/http"

// GET /is_login cookieから受け取った sessionID を元にユーザーがログインしているかどうか確認
func FindUserBySession(w http.ResponseWriter, r *http.Request) {

}
