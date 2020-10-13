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
	Type        string
	ImageBase64 string
}

func NewClothUseCase(clothRepo repository.Cloth) *ClothUseCase {
	return &ClothUseCase{clothRepo: clothRepo}
}

// 服を新規追加
func (cu *ClothUseCase) CreateCloth(cloth *Cloth) (int, error) {
	var clothEntityModel entity.Cloth
	clothEntityModel.Name = cloth.Name
	clothEntityModel.Price = cloth.Price
	clothEntityModel.Type = cloth.Type
	if cloth.ImageBase64 != "" {
		clothEntityModel.ImageUrl = service.UploadS3(cloth.ImageBase64)
	} else {
		clothEntityModel.ImageUrl = "no_url"
	}

	// MySQL に服情報を追加
	clothID, err := cu.clothRepo.Create(&clothEntityModel)
	if err != nil {
		fmt.Print(err)
		return 0, err
	}

	return clothID, nil
}

// すべての服情報を取得
func (cu *ClothUseCase) GetAll() *[]entity.Cloth {
	clothes := cu.clothRepo.GetAll()
	return clothes
}
