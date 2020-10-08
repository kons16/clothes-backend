package Redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"os"
)

func NewRedisDB() (*redis.Client, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, // use default DB
	})
	return rdb, nil
}
