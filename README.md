# clothes-backend
標準ライブラリで作る, DDD構成のAPIサーバー. ([フロントエンドのリポジトリ](https://github.com/sunakane/team7-mock)
)  
[構成の参考にしたドキュメント](https://github.com/camphor-/relaym-server/blob/master/docs/application_architecture.md)  

APIの仕様は [doc/api.md](https://github.com/kons16/team7-backend/blob/master/doc/api.md) を参考にしてください.   

<img src="https://user-images.githubusercontent.com/31591102/103457722-14ae9500-4d45-11eb-8d82-a4628fd4791a.png" height="350">
  
## 環境構築
```
$ git clone https://github.com/kons16/clothes-backend.git
$ cd clothes-backend
$ docker-compose up
```
DBの起動とS3へのアクセスには `.env` を設定する必要があります.

## テスト
```
$ go test -v ./...
```  

## データベース
### MySQL
ユーザー情報(usersテーブル), 服情報(clothesテーブル), コーディネート情報(cordinatesテーブル) は MySQL に保存しています.
```
$ docker exec -it [db_container_id] /bin/bash

# mysql -u user -p

> USE [database_name];
> SHOW TABLES;
```
マイグレーションには sql-migrate を使用しています.  
```
$ docker exec -it [api_container_id] sh

# sql-migrate new table_name

# sql-migrate up
```

### Redis
ユーザーの SessionID の保存には Redis を使用しています.  
Redis には key が SessionID, field が UserID と ExpiresAt をhashで保存しています.  
```
$ docker exec -it [redis_container_id] sh

# redis-cli

> keys *
1) "xxxyyyzzz"

> hgetall xxxyyyzzz
1) "UserID"
2) "112233"
3) "ExpiresAt"
4) "2020-10-11 05:11:32"
```
