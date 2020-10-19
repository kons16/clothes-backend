# API仕様
エンドポイントは `localhost:8000/api/v1/` です。  
データ形式は `application/json` です。  

ブラウザで確認するときは [Talend API Tester](https://chrome.google.com/webstore/detail/talend-api-tester-free-ed/aejoelaoggembcahagimdiliamlcdmfm) 等を利用してください。


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
`sex` が `0` のとき男性、`1` のとき女性
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
## 服の新規追加
`POST /cloth`  
#### リクエスト
`enctype="multipart/form-data"` を指定する必要があります。  
アップロードする `image` は `base64`です。  
`type` が `a` のときはトップ、`b` のときはダウン。
```
{
  "name": "乃木坂Tシャツ",
  "price": "3000",
  "image": "xxxxyyyyyzzzz",
  "type": "a"
}
```
#### レスポンス
アップロードされた服の id を返します。
```
{
  "cloth_id": "123"
}
```

## 服の全取得
`GET /get_cloth`  
#### リクエスト
なし (今後 query で検索できるようにするかも)
#### レスポンス
```
{
  "clothes":[
    {"ID": "123", "ImageUrl": "https:~~~.jpeg", "Name": "乃木坂Tシャツ", "Price": "3000", "type": "a"},
    {"ID": "124", "ImageUrl": "no_url", "Name": "ベージュパンツ", "Price": "5000", "type": "b"}
  ]
}
```

## 服の購入
`POST /buy`  
#### リクエスト
cookieをHeaderに付与
```
{
  "cloth_id": "123"
}
```
#### レスポンス
```
{
  "message": "success"
}
```

## 購入した服の取得
`GET /my_cloth`  
#### リクエスト
cookieをHeaderに付与
#### レスポンス
```
{
  "clothes":[
    {"ID": "123", "ImageUrl": "https:~~~.jpeg", "Name": "乃木坂Tシャツ", "Price": "3000", "type": "a"},
    {"ID": "124", "ImageUrl": "no_url", "Name": "ベージュパンツ", "Price": "5000", "type": "b"}
  ]
}
```

## コーディネート関連
## コーディネートの新規追加
`POST /cordinate`  
#### リクエスト
cookie を Header に付与。
```
{
  "title": "マイコーデ１",
  "top_cloth_id": "123",
  "pant_cloth_id": "456"
}
```
#### レスポンス
```
{
  "message": "success"
}
```