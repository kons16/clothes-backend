package Redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/kons16/team7-backend/domain/entity"
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
		"SessionID": sessionID,
		"ExpiresAt": expiresAt.Format(layout),
	}

	// key: UserID, Value: SessionID, ExpiresAt
	sr.rdMap.HMSet(ctx, string(userID), m)

	return nil
}
