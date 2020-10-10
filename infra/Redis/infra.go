package Redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

func NewRedisDB() (*redis.Client, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	return rdb, nil
}
