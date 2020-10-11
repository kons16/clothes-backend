package repository

import "github.com/kons16/team7-backend/domain/entity"

type Cloth interface {
	Create(cloth *entity.Clothe) (int, error)
}
