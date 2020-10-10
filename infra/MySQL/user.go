package MySQL

import (
	"github.com/jmoiron/sqlx"
	"github.com/kons16/team7-backend/domain/entity"
	"time"
)

// UserRepository は repository.UserRepository を満たす構造体
type UserRepository struct {
	dbMap *sqlx.DB
}

func NewUserRepository(dbMap *sqlx.DB) *UserRepository {
	return &UserRepository{dbMap: dbMap}
}

// ユニークなIDを生成
func (r *UserRepository) generateID() (int, error) {
	var id int
	err := r.dbMap.Get(&id, "SELECT UUID_SHORT()")
	return id, err
}

// FindByID は指定されたIDを持つユーザをDBから取得します
func (r *UserRepository) FindByID(id string) (*entity.User, error) {
	return nil, nil
}

// Create はユーザを新規保存します
func (r *UserRepository) Create(user *entity.User) (int, error) {
	now := time.Now()
	id, _ := r.generateID()

	_, err := r.dbMap.Exec(
		`INSERT INTO users
			(id, name, submit_id, year, sex, password_hash, created_at, updated_at)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		id, user.Name, user.SubmitID, user.Year, user.Sex, user.PasswordHash, now, now,
	)
	return id, err
}
