package usecase

import (
	"fmt"
	"github.com/kons16/team7-backend/domain/entity"
	"github.com/kons16/team7-backend/domain/repository"
	"github.com/kons16/team7-backend/domain/service"
)

type ClothUseCase struct {
	clothRepo repository.Cloth
}

// 新規登録する服の構造体 (まだbase64)
type Cloth struct {
	Name        string
	Price       string
	ImageBase64 string
}

func NewClothUseCase(clothRepo repository.Cloth) *ClothUseCase {
	return &ClothUseCase{clothRepo: clothRepo}
}

func (cu *ClothUseCase) CreateCloth(cloth *Cloth) (int, error) {
	var clothEntityModel entity.Cloth
	clothEntityModel.Name = cloth.Name
	clothEntityModel.Price = cloth.Price
	clothEntityModel.ImageUrl = service.UploadS3(cloth.ImageBase64)

	// MySQL に服情報を追加
	clothID, err := cu.clothRepo.Create(&clothEntityModel)
	if err != nil {
		fmt.Print(err)
		return 0, err
	}

	return clothID, nil
}
