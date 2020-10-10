# API仕様
エンドポイントは `localhost:8000/api/v1/` です。  
データ形式は　`application/json` です。

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
- `POST /user`  
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

## 服関連

