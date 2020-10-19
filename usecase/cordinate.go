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

// コーディネートを新規追加
func (cdu *CordinateUseCase) CreateCordinate(cordinate *entity.Cordinate) error {
	return nil
}

// すべてのコーディネート情報を取得
func (cdu *CordinateUseCase) GetAll(userID int) *[]entity.Cordinate {
	return nil
}
