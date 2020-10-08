package MySQL

import (
	"github.com/jmoiron/sqlx"
	"github.com/kons16/team7-backend/domain/entity"
)

// UserRepository は repository.UserRepository を満たす構造体
type UserRepository struct {
	dbMap *sqlx.DB
}

func NewUserRepository(dbMap *sqlx.DB) *UserRepository {
	return &UserRepository{dbMap: dbMap}
}

// FindByID は指定されたIDを持つユーザをDBから取得します
func (r *UserRepository) FindByID(id string) (*entity.User, error) {
	return nil, nil
}

// Create はユーザを新規保存します。
func (r *UserRepository) Create(user *entity.User) error {
	return nil
}
