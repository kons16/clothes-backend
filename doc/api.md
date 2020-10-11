# API仕様
エンドポイントは `localhost:8000/api/v1/` です。  
データ形式は `application/json` です。

## サーバー起動テスト
- `GET /hello`  
#### リクエスト
空
#### レスポンス
```
{
  "message": "hello"
}
```

## ユーザー登録関連

## 新規登録
`POST /user`
  
#### リクエスト
```
{
  "name": "kono",
  "submit_id": "kono1997",
  "password": "password",
  "year": "22",
  "sex": "0"
}
```
#### レスポンス
```
{
  "sessionID": "xxxyyyzzz"
}
```

## ログイン
`POST /login`  
#### リクエスト
```
{
  "submit_id": "kono1997",
  "password": "password"
}
```
#### レスポンス
```
{
  "sessionID": "xxxyyyzzz"
}
```

## ログインしているかどうかの確認
`GET /is_login`  
#### リクエスト
cookieをHeaderに付与
#### レスポンス
`isLogin` が `true` のときログインしている。  
`false` のときログインしていない。
```
{
  "isLogin": "true"
}
```

## ログアウト
`GET /logout`  
#### リクエスト
cookieをHeaderに付与
#### レスポンス
`logout` が `true` のときログアウト成功。  
`false` のときログアウト失敗。
```
{
  "logout": "true"
}
```

## 服関連

