package MySQL

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kons16/team7-backend/domain/entity"
	"time"
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

// Create はコーディネートを新規で追加します
func (r *CordinateRepository) Create(cordinate *entity.Cordinate) error {
	now := time.Now()
	id, _ := r.generateID()

	_, err := r.dbMap.Exec(
		`INSERT INTO cordinates
			(id, title, top_cloth_id, pant_cloth_id, user_id, created_at, updated_at)
			VALUES (?, ?, ?, ?, ?, ?, ?)`,
		id, cordinate.Title, cordinate.TopClothID, cordinate.PantClothID, cordinate.UserID, now, now,
	)
	return err
}

// Get は ユーザーに紐づく全てのコーディネート情報を取得します
func (r *CordinateRepository) Get(userID int) *[]entity.Cordinate {
	var cordinates []entity.Cordinate
	err := r.dbMap.Select(&cordinates, `SELECT id, title, top_cloth_id, pant_cloth_id FROM cordinates WHERE user_id = ?`, userID)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &cordinates
}
