# team7-backend
標準ライブラリで作る、DDD構成のAPIサーバー.  
[構成の参考にしたドキュメント](https://github.com/camphor-/relaym-server/blob/master/docs/application_architecture.md)  

APIの仕様は [doc/api.md](https://github.com/kons16/team7-backend/blob/master/doc/api.md) を参考にしてください.
  
## 環境構築
```
$ git clone https://github.com/kons16/team7-backend.git
$ cd team7-backend
$ docker-compose up
```

## Redis
ユーザーの SessionID の保存には Redis を使用しています.  
Redis には key が SessionID, field が UserID と ExpiresAt をhashで保存しています.  
```
$ docker exec -it [container_id] sh

# redis-cli

> keys *
1) "xxxyyyzzz"

> hgetall xxxyyyzzz
1) "UserID"
2) "112233"
3) "ExpiresAt"
4) "2020-10-11 05:11:32"
```
