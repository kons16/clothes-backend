package usecase

import (
	"fmt"
	"github.com/kons16/team7-backend/domain/entity"
	"github.com/kons16/team7-backend/domain/repository"
)

type ClothUseCase struct {
	clothRepo repository.Cloth
}

func NewClothUseCase(clothRepo repository.Cloth) *ClothUseCase {
	return &ClothUseCase{clothRepo: clothRepo}
}

func (cu *ClothUseCase) CreateCloth(cloth *entity.Clothe) (int, error) {
	// MySQL に服情報を追加
	clothID, err := cu.clothRepo.Create(cloth)
	if err != nil {
		fmt.Print(err)
		return 0, err
	}

	return clothID, nil
}
