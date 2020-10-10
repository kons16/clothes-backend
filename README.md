# team7-backend
標準ライブラリで作る、DDD構成のAPIサーバー.  
[構成の参考にしたドキュメント](https://github.com/camphor-/relaym-server/blob/master/docs/application_architecture.md)  
  
## 環境構築
```
$ git clone https://github.com/kons16/team7-backend.git
$ cd team7-backend
$ docker-compose up
```

## Redis
ユーザーの SessionID の保存には Redis を使用しています.  
Redis には key が UserID, value が SessionID と ExpiresAt を保存しています.  
```
$ docker exec -it [container_id] sh

# redis-cli

> keys *
1) "111222333"

> hgetall 111222333
1) "SessionID"
2) "xxxyyyzzz"
3) "ExpiresAt"
4) "2020-10-11 05:11:32"
```
