package Redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/kons16/team7-backend/domain/entity"
	"strconv"
	"time"
)

type SessionRepository struct {
	rdMap *redis.Client
}

func NewSessionRepository(rdMap *redis.Client) *SessionRepository {
	return &SessionRepository{rdMap: rdMap}
}

func (sr *SessionRepository) FindUserBySession(sessionID int) (*entity.User, error) {
	return nil, nil
}

// UserIDに紐づくSessionIDをRedisに保存
func (sr *SessionRepository) CreateUserSession(userID int, sessionID string) error {
	ctx := context.Background()
	expiresAt := time.Now().Add(24 * time.Hour)

	layout := "2006-01-02 15:04:05"

	m := map[string]string{
		"UserID":    strconv.Itoa(userID),
		"ExpiresAt": expiresAt.Format(layout),
	}

	// key: SessionID, Hash [ UserID, ExpiresAt ]
	for field, val := range m {
		fmt.Println("Inserting", "field:", field, "val:", val)
		err := sr.rdMap.HSet(ctx, sessionID, field, val).Err()
		if err != nil {
			fmt.Println("sr.rdMap.HSet Error:", err)
		}
	}

	/*
		h, err := sr.rdMap.HGet(ctx, key, "SessionID").Result()
		if err != nil {
			fmt.Println("redis.Client.HGet Error:", err)
		}
		fmt.Println(h)
	*/

	return nil
}
