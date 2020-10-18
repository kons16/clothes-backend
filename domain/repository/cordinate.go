package repository

import "github.com/kons16/team7-backend/domain/entity"

type Cordinate interface {
	Create(cloth *entity.Cordinate) *entity.Cordinate
	Get(userID int) *[]entity.Cordinate
}
