package usecase

import (
	"github.com/kons16/team7-backend/domain/entity"
	"github.com/kons16/team7-backend/domain/repository"
)

type CordinateUseCase struct {
	cordiRepo repository.Cordinate
}

func NewCordinateUseCase(cordiRepo repository.Cordinate) *CordinateUseCase {
	return &CordinateUseCase{cordiRepo: cordiRepo}
}

// 服を新規追加
func (cdu *CordinateUseCase) CreateCloth(cloth *entity.Cordinate) *entity.Cordinate {
	return nil
}

// すべての服情報を取得
func (cdu *CordinateUseCase) GetAll(userID int) *[]entity.Cordinate {
	return nil
}
