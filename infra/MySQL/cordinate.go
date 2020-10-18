package MySQL

import (
	"github.com/jmoiron/sqlx"
	"github.com/kons16/team7-backend/domain/entity"
)

// CordinateRepository は repository.CordinateRepository を満たす構造体
type CordinateRepository struct {
	dbMap *sqlx.DB
}

func NewCordinateRepository(dbMap *sqlx.DB) *CordinateRepository {
	return &CordinateRepository{dbMap: dbMap}
}

// ユニークなIDを生成
func (r *CordinateRepository) generateID() (int, error) {
	var id int
	err := r.dbMap.Get(&id, "SELECT UUID_SHORT()")
	return id, err
}

// Create は服を新規で追加します
func (r *CordinateRepository) Create(cloth *entity.Cordinate) *entity.Cordinate {
	return nil
}

// GetAll は すべての服情報を取得します
func (r *CordinateRepository) Get(userID int) *[]entity.Cordinate {
	return nil
}
