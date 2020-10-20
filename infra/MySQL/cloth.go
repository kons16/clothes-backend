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

// GetBuyCloth は ユーザーが購入した服情報を返す
func (r *ClothRepository) GetBuyCloth(userID int) *[]entity.Cloth {
	var userCloths []interface{}
	err := r.dbMap.Select(&userCloths, `SELECT cloth_id FROM user_clothes WHERE user_id = ?`, userID)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	q := `SELECT id, name, price, image_url, type FROM clothes WHERE id IN (?)`
	sql, params, err := sqlx.In(q, userCloths)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var cloths []entity.Cloth
	// fmt.Println(sql)   SELECT id, name, price, image_url, type FROM clothes WHERE id IN (?, ?)
	// fmt.Println(params)   [{98947064062279683} {98947064062279684}]

	err = r.dbMap.Select(&cloths, sql, params...)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println(&cloths)
	return &cloths
}

// GetByIDs は 服ID のすべての服情報を返す
func (r *ClothRepository) GetByIDs(clothIDs []int) *[]entity.Cloth {
	q := `SELECT id, name, price, image_url, type FROM clothes WHERE id IN (?)`
	sql, params, err := sqlx.In(q, clothIDs)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var cloths []entity.Cloth
	// fmt.Println(sql)   SELECT id, name, price, image_url, type FROM clothes WHERE id IN (?, ?)
	// fmt.Println(params)   [{98947064062279683} {98947064062279684}]

	err = r.dbMap.Select(&cloths, sql, params...)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println(&cloths)
	return &cloths
}
