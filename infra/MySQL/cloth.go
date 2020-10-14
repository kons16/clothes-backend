package MySQL

import (
	"fmt"
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
func (r *ClothRepository) Create(cloth *entity.Cloth) (int, error) {
	now := time.Now()
	id, _ := r.generateID()

	_, err := r.dbMap.Exec(
		`INSERT INTO clothes
			(id, name, price, image_url, type, created_at, updated_at)
			VALUES (?, ?, ?, ?, ?, ?, ?)`,
		id, cloth.Name, cloth.Price, cloth.ImageUrl, cloth.Type, now, now,
	)
	return id, err
}

// GetAll は すべての服情報を取得します
func (r *ClothRepository) GetAll() *[]entity.Cloth {
	var cloth []entity.Cloth
	err := r.dbMap.Select(&cloth, `SELECT id, name, price, image_url, type FROM clothes`)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &cloth
}

// CreateUserCloth は　user_clothesテーブル に userID と clothID を追加します
func (r *ClothRepository) CreateUserCloth(userID int, clothID int) error {
	now := time.Now()
	id, _ := r.generateID()

	_, err := r.dbMap.Exec(
		`INSERT INTO user_clothes
			(id, user_id, cloth_id, created_at, updated_at)
			VALUES (?, ?, ?, ?, ?)`,
		id, userID, clothID, now, now,
	)
	return err
}
