package Redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/kons16/team7-backend/domain/entity"
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

func (sr *SessionRepository) CreateUserSession(id int) (string, error) {

	return "", nil
}
