package MySQL

import (
	"github.com/jmoiron/sqlx"
	"github.com/kons16/team7-backend/domain/entity"
	"time"
)

// ClothRepository は repository.ClothRepository を満たす構造体
type ClothRepository struct {
	dbMap *sqlx.DB
}

func NewClothRepository(dbMap *sqlx.DB) *ClothRepository {
	return &ClothRepository{dbMap: dbMap}
}

// ユニークなIDを生成
func (r *ClothRepository) generateID() (int, error) {
	var id int
	err := r.dbMap.Get(&id, "SELECT UUID_SHORT()")
	return id, err
}

// Create は服を新規で追加します
func (r *ClothRepository) Create(cloth *entity.Clothe) (int, error) {
	now := time.Now()
	id, _ := r.generateID()

	_, err := r.dbMap.Exec(
		`INSERT INTO clothes
			(id, name, price, url, created_at, updated_at)
			VALUES (?, ?, ?, ?, ?, ?)`,
		id, cloth.Name, cloth.Price, cloth.ImageUrl, now, now,
	)
	return id, err
}
