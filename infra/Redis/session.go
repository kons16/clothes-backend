package Redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

type SessionRepository struct {
	rdMap *redis.Client
}

func NewSessionRepository(rdMap *redis.Client) *SessionRepository {
	return &SessionRepository{rdMap: rdMap}
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

	return nil
}

// SessionID に紐づく UserID があるかどうか確認
func (sr *SessionRepository) CheckBySession(sessionID string) int {
	ctx := context.Background()
	// TODO: 時間が切れてないか確認する now := time.Now()

	getUserID, err := sr.rdMap.HGet(ctx, sessionID, "UserID").Result()
	if err != nil {
		fmt.Println("redis.Client.HGet Error:", err)
		return 0
	}

	// Redis には string で格納されているので int に変換
	getUserIntID, _ := strconv.Atoi(getUserID)
	return getUserIntID
}

// Logout は SessionID のカラムを Redis から削除する
func (sr *SessionRepository) Logout(sessionID string) bool {
	ctx := context.Background()

	err := sr.rdMap.Del(ctx, sessionID).Err()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
