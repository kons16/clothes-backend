# API仕様
エンドポイントは `localhost:8000/api/v1/` です。  
データ形式は `application/json` です。

## サーバーが起動しているかテスト
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
  "session_id": "xxxyyyzzz"
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
  "session_id": "xxxyyyzzz"
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
  "is_login": "true"
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
  "is_logout": "true"
}
```

## 服関連
`POST /cloth`  
#### リクエスト
`enctype="multipart/form-data"` を指定する必要があります。  
アップロードする `image` は `base64`です。
```
{
  "name": "乃木坂Tシャツ",
  "price": "3000",
  "image": "xxxxyyyyyzzzz"
}
```
#### レスポンス
アップロードされた画像の id を返します。
```
{
  "image_id": "123"
}
```